package main

import (
	"flag"
	"fmt"
	"os"

	"mock-data-generator/cmd"
)

func main() {
	var numRecords int
	flag.IntVar(&numRecords, "numRecords", 100, "Number of records to generate")
	flag.Parse()

	if err := cmd.Run(numRecords); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
