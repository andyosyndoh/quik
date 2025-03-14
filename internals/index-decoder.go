package internals

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

type entry struct {
	simhash uint64
	offsets []int64
}

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

	entries := make(chan entry, len(indexData.Index))
	outputChan := make(chan string, runtime.NumCPU()*2)
	return nil
}
