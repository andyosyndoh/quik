package internals

import (
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
)

// RunIndex processes an input file to build an index and serialize it to an output file.
// It performs the following steps:
// 1. Validates the chunk size to ensure it is greater than 0.
// 2. Validates the input file using the ValidateInputFile function.
// 3. Builds the index using the NewFileIndex and BuildIndex functions.
// 4. Prepares the IndexData structure with the file name, chunk size, and index data.
// 5. Starts a goroutine to process the index data concurrently using the IndexFileDecoder function.
// 6. Serializes the index data to the specified output file using gob encoding.
//
// Parameters:
// - inputFile: The path to the input file to be indexed.
// - chunkSize: The size of each chunk for indexing.
// - outputFile: The path to the output file where the serialized index data will be saved.
//
// Returns:
// - error: An error if any step fails, otherwise nil.
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
