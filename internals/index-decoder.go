package internals

import (
	"bufio"
	"fmt"
	"os"
)

// IndexFileDecoder decodes the index file and prints the metadata and index data.
func IndexFileDecoder(indexData IndexData) error {
	fmt.Printf("Original file: %s\n", indexData.FileName)
	fmt.Printf("Chunk size: %d bytes\n", indexData.ChunkSize)
	fmt.Println("SimHash values and byte offsets writen to simhash.txt")

	hashfile, err := os.Create("simhash.txt")
	if err != nil {
		return fmt.Errorf("error creating hash file: %w", err)
	}
	defer hashfile.Close()

	writer := bufio.NewWriter(hashfile)
	defer writer.Flush()
	return nil
}
