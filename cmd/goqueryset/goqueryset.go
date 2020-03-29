package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/to6ka/go-queryset/queryset"
)

var outFile = flag.String("out", "{in}_queryset.go", "path to output file")

func main() {
	flag.Parse()

	inFile := os.Getenv("GOFILE")
	inFile = path.Base(inFile)

	fileStat, err := os.Stat(inFile)
	if err != nil {
		log.Fatalf("failed to access file %s", inFile)
	}
	if !fileStat.Mode().IsRegular() {
		log.Fatalf("file %s is not a regular file", inFile)
	}

	if strings.Contains(*outFile, "{in}") {
		inName := strings.TrimSuffix(inFile, ".go")
		*outFile = strings.Replace(*outFile, "{in}", inName, 1)
	}

	if err := queryset.GenerateQuerySets(inFile, *outFile); err != nil {
		log.Fatalf("can't generate query sets: %s", err)
	}
}
