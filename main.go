package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"textindexer/internals"
)

// main is the entry point of the application. It parses command-line arguments
// and executes the appropriate command based on the provided options.
//
// Usage: textindex [options]
//
// Commands:
//
//	-c index  : Indexes the input file and generates an output index file.
//	  Options:
//	    -i string : Input file path (required)
//	    -s int    : Chunk size in bytes (default: 4096)
//	    -o string : Output index file path (required)
//
//	-c lookup : Looks up a SimHash value in the specified index file.
//	  Options:
//	    -i string : Index file path (required)
//	    -h string : SimHash value to lookup (required)
//
//	-c fuzzy  : Performs a fuzzy search for a SimHash value in the specified index file.
//	  Options:
//	    -i string : Index file path (required)
//	    -h string : SimHash value for fuzzy search (required)
//
// If an unknown command or invalid options are provided, the program will print an error message and exit.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: textindex [options]")
		os.Exit(1)
	}

	command := ""
	if os.Args[1] == "-c" {
		command = os.Args[2]
	} else {
		fmt.Println("Unknown command option")
		os.Exit(1)
	}
	args := os.Args[3:]

	switch command {
	case "index":
		indexFlags := flag.NewFlagSet("index", flag.ExitOnError)
		inputFile := indexFlags.String("i", "", "Input file path (required)")
		chunkSize := indexFlags.Int("s", 4096, "Chunk size in bytes")
		outputFile := indexFlags.String("o", "", "Output index file path (required)")
		indexFlags.Parse(args)

		if *inputFile == "" || *outputFile == "" {
			fmt.Println("Error: -i and -o are required for index command")
			os.Exit(1)
		}

		if !strings.HasSuffix(*outputFile, ".idx") {
			fmt.Println("Error: please provide an index file (with .idx extension)")
			os.Exit(1)
		}

		if *chunkSize <= 0 {
			fmt.Println("Error: invalid chunk size")
			os.Exit(1)
		}

		if err := internals.RunIndex(*inputFile, *chunkSize, *outputFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "lookup":
		lookupFlags := flag.NewFlagSet("lookup", flag.ExitOnError)
		indexFile := lookupFlags.String("i", "", "Index file path")
		simHashStr := lookupFlags.String("h", "", "SimHash value to lookup")
		lookupFlags.Parse(args)

		if *indexFile == "" || *simHashStr == "" {
			fmt.Println("Error: -i and -h are required for lookup command")
			os.Exit(1)
		}

		if !strings.HasSuffix(*indexFile, ".idx") {
			fmt.Println("Error: please input an index file")
			os.Exit(1)
		}

		if err := internals.RunLookup(*indexFile, *simHashStr); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "fuzzy":
		// New fuzzy command logic
		fuzzyFlags := flag.NewFlagSet("fuzzy", flag.ExitOnError)
		indexFile := fuzzyFlags.String("i", "", "Index file path")
		simHashStr := fuzzyFlags.String("h", "", "SimHash value for fuzzy search")
		fuzzyFlags.Parse(args)

		if *indexFile == "" || *simHashStr == "" {
			fmt.Println("Error: -i and -h are required for fuzzy command")
			os.Exit(1)
		}

		if !strings.HasSuffix(*indexFile, ".idx") {
			fmt.Println("Error: please input an index file")
			os.Exit(1)
		}

		if err := internals.RunFuzzy(*indexFile, *simHashStr); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid command. Use 'index' or 'lookup'.")
		os.Exit(1)

	}
}
