package internals

import (
	"fmt"
)

// IndexFileDecoder decodes the index file and prints the metadata and index data.
func IndexFileDecoder(indexData IndexData) error {
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
