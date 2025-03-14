package internals

import (
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
)

// RunIndex builds an index for the input file, validates the input, and serializes it to the output file.
// It also starts a goroutine to process the index data concurrently using IndexFileDecoder.
func RunIndex(inputFile string, chunkSize int, outputFile string) error {
	// Ensures chunkSize is valid to prevent infinite loops or excessive resource usage.
	if chunkSize <= 0 {
		return fmt.Errorf("invalid chunk size: %d, must be greater than 0", chunkSize)
	}
	// Validate input file
	if err := ValidateInputFile(inputFile); err != nil {
		return err
	}

	// Build the index
	fi := NewFileIndex(chunkSize, runtime.NumCPU())
	if err := fi.BuildIndex(inputFile); err != nil {
		return fmt.Errorf("error building index: %v", err)
	}

	// Prepare the IndexData structure
	indexData := IndexData{
		FileName:  inputFile,
		ChunkSize: chunkSize,
		Index:     fi.index.m,
	}

	// Start a goroutine to process the index data concurrently
	go func(data IndexData) {
		if err := IndexFileDecoder(data); err != nil {
			fmt.Printf("error in IndexFileDecoder: %v\n", err)
		}
	}(indexData)

	// Serialize the index data to the output file
	dataFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating index file: %v", err)
	}
	defer dataFile.Close()

	encoder := gob.NewEncoder(dataFile)
	if err := encoder.Encode(indexData); err != nil {
		return fmt.Errorf("error encoding index data: %v", err)
	}

	return nil
}
