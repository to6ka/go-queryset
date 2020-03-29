package main

import (
	"flag"
	"log"
	"path"
	"strings"

	"github.com/to6ka/go-queryset/queryset"
)

func main() {
	inFile := flag.String("in", "models.go", "path to input file")
	outFile := flag.String("out", "{in}_queryset.go", "path to output file")
	flag.Parse()

	if strings.Contains(*outFile, "{in}") {
		baseName := path.Base(*inFile)
		*outFile = strings.Replace(*outFile, "{in}", baseName, 1)
	}

	if err := queryset.GenerateQuerySets(*inFile, *outFile); err != nil {
		log.Fatalf("can't generate query sets: %s", err)
	}
}
