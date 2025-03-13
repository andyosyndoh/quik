package main

import (
	"flag"
	"fmt"
	"os"
)

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

		if *chunkSize <= 0 {
			fmt.Println("Error: invalid chunk size")
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

	default:
		fmt.Println("Invalid command. Use 'index' or 'lookup'.")
		os.Exit(1)

	}
}
