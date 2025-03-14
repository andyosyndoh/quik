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
	fmt.Println("SimHash values and byte offsets:")

	hashfile, err := os.Create("simhash.txt")
	if err != nil {
		return fmt.Errorf("error creating hash file: %w", err)
	}
	defer hashfile.Close()

	writer := bufio.NewWriter(hashfile)
	defer writer.Flush()

	numWorkers := runtime.NumCPU()

	entries := make(chan entry, len(indexData.Index))
	outputChan := make(chan string, numWorkers*2)

	// Send entries to be processed
	go func() {
		for simhash, offsets := range indexData.Index {
			entries <- entry{simhash, offsets}
		}
		close(entries)
	}()

	var wg sync.WaitGroup

	for simhash, offsets := range indexData.Index {
		fmt.Printf("SimHash: %x\n", simhash)
		for _, offset := range offsets {
			fmt.Printf("  Byte offset: %d\n", offset)
		}
	}
	return nil
}
