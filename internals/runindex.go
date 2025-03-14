package internals

import (
	"fmt"
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
	return nil
}