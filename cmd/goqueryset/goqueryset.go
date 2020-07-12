package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/to6ka/go-queryset/queryset"
)

var paramOut = flag.String("out", "{in}_queryset.go", "path to output file")
var paramDBType = flag.String("dbtype", "*github.com/jinzhu/gorm.DB", "db type to use for generation")

func main() {
	flag.Parse()

	inFile := os.Getenv("GOFILE")
	inFile = path.Base(inFile)

	outFile := *paramOut
	if strings.Contains(outFile, "{in}") {
		inName := strings.TrimSuffix(inFile, ".go")
		outFile = strings.Replace(outFile, "{in}", inName, 1)
	}

	dbType, dbImport, err := parseDBType(*paramDBType)
	if err != nil {
		log.Fatalf("failed to parse db type: %s", err)
	}

	fileStat, err := os.Stat(inFile)
	if err != nil {
		log.Fatalf("failed to access file %s: %s", inFile, err)
	}
	if !fileStat.Mode().IsRegular() {
		log.Fatalf("file %s is not a regular file", inFile)
	}

	log.Printf("generating goqueryset in=%s out=%s dbtype=%s dbimport=%s", inFile, outFile, dbType, dbImport)
	err = queryset.GenerateQuerySets(inFile, outFile, queryset.Config{
		DBType:   dbType,
		DBImport: dbImport,
	})
	if err != nil {
		log.Fatalf("can't generate query sets: %s", err)
	}
}

func parseDBType(typeLine string) (dbType string, dbImport string, err error) {
	s := typeLine
	if s[0] == '*' {
		dbType = "*"
		s = s[1:]
	}

	rawType := path.Base(s)
	dbType += rawType

	typeParts := strings.Split(rawType, ".")
	if len(typeParts) != 2 {
		return "", "", fmt.Errorf("one dot in db type expected (original:%q, got:%q)", typeLine, dbType)
	}

	pkg := typeParts[0]
	dbImport = path.Dir(s) + "/" + pkg

	return dbType, dbImport, nil
}
