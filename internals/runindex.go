package internals

import (
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
)

func runIndex(inputFile string, chunkSize int, outputFile string) error {
	fi := NewFileIndex(chunkSize, runtime.NumCPU())
	if err := fi.BuildIndex(inputFile); err != nil {
		return fmt.Errorf("error building index: %v", err)
	}

	indexData := IndexData{
		FileName:  inputFile,
		ChunkSize: chunkSize,
		Index:     fi.index.m,
	}

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
