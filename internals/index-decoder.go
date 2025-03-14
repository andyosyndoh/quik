package internals

import (
	"encoding/gob"
	"fmt"
	"os"
)

// IndexFileDecoder decodes the index file and prints the metadata and index data.
func IndexFileDecoder(file string) error {
	// Load index file into memory
	dataFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error opening index file: %w", err)
	}
	defer dataFile.Close()

	// Decode the file content into IndexData struct
	var indexData IndexData
	if err := gob.NewDecoder(dataFile).Decode(&indexData); err != nil {
		return fmt.Errorf("error decoding index data: %v\n", err)
	}

	fmt.Printf("Original file: %s\n", indexData.FileName)
	fmt.Printf("Chunk size: %d bytes\n", indexData.ChunkSize)
	fmt.Println("SimHash values and byte offsets:")

	
	for simhash, offsets := range indexData.Index {
		fmt.Printf("SimHash: %x\n", simhash)
		for _, offset := range offsets {
			fmt.Printf("  Byte offset: %d\n", offset)
		}
	}
	return nil
}
