package internals

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

type entry struct {
	simhash uint64
	offsets []int64
}

// IndexFileDecoder processes the given IndexData and writes SimHash values and their byte offsets to a file named "simhash.txt".
// It utilizes a worker pool to handle the processing concurrently.
//
// Parameters:
//   - indexData: The IndexData containing the file name, chunk size, and index map of SimHash values to byte offsets.
//
// Returns:
//   - error: An error if any occurs during file creation or writing, otherwise nil.
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

	// Send entries to be processed
	go func() {
		for simhash, offsets := range indexData.Index {
			entries <- entry{simhash, offsets}
		}
		close(entries)
	}()

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for e := range entries {
				var buf strings.Builder
				fmt.Fprintf(&buf, "SimHash: %x\n", e.simhash)
				for _, offset := range e.offsets {
					fmt.Fprintf(&buf, "  Byte offset: %d\n", offset)
				}
				outputChan <- buf.String()
			}
		}()
	}

	// Close output channel when all workers finish
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// Write outputs to file
	for output := range outputChan {
		if _, err := writer.WriteString(output); err != nil {
			return fmt.Errorf("error writing to hash file: %w", err)
		}
	}

	return nil
}
