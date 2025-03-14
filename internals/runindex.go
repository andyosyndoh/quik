package internals

import (
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
)

// runIndex builds an index for the input file and serializes it to the output file.
func RunIndex(inputFile string, chunkSize int, outputFile string) error {
	fi := NewFileIndex(chunkSize, runtime.NumCPU())
	if err := fi.BuildIndex(inputFile); err != nil {
		return fmt.Errorf("error building index: %v", err)
	}
    
	// Prepare the IndexData structure to store metadata and the index itself
	indexData := IndexData{
		FileName:  inputFile,
		ChunkSize: chunkSize,
		Index:     fi.index.m,
	}

	go func(indexdata IndexData) {
		IndexFileDecoder(indexdata)
	}(indexData)

	// Create the output file to store the serialized index
	dataFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating index file: %v", err)
	}
	defer dataFile.Close()

	// Encode the IndexData structure using gob and write it to the output file
	encoder := gob.NewEncoder(dataFile)
	if err := encoder.Encode(indexData); err != nil {
		return fmt.Errorf("error encoding index data: %v", err)
	}
	return nil
}
