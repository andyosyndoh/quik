package internals

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// RunFuzzy performs a fuzzy search for SimHashes with a Hamming distance of 1 or 2.
func RunFuzzy(indexFile, simHashStr string) error {
	dataFile, err := os.Open(indexFile)
	if err != nil {
		return fmt.Errorf("error opening index file: %v", err)
	}
	defer dataFile.Close()

	// Decode the index data
	var indexData IndexData
	decoder := gob.NewDecoder(dataFile)
	if err := decoder.Decode(&indexData); err != nil {
		return fmt.Errorf("error decoding index data: %v", err)
	}

	// Check if the original file exists
	if _, err := os.Stat(indexData.FileName); os.IsNotExist(err) {
		return fmt.Errorf("original file %s not found", indexData.FileName)
	}

	// Parse the provided SimHash
	simHash, err := strconv.ParseUint(simHashStr, 16, 64)
	if err != nil {
		return fmt.Errorf("invalid SimHash value: %v", err)
	}

	// Open the original file
	file, err := os.Open(indexData.FileName)
	if err != nil {
		return fmt.Errorf("error opening original file: %v", err)
	}
	defer file.Close()

	for hash, offsets := range indexData.Index {
		distance := hammingdistance(simHash, hash)
		if distance == 1 {
			for _, offset := range offsets {
				chunk := make([]byte, indexData.ChunkSize)
				n, err := file.ReadAt(chunk, offset)
				if err != nil && err != io.EOF {
					return fmt.Errorf("error reading chunk at offset %d: %v", offset, err)
				}
				chunk = chunk[:n]

				// Extract a phrase from the chunk
				words := strings.Fields(string(chunk))
				phrase := strings.Join(words, " ")
				if len(phrase) > 50 {
					phrase = phrase[:50] + "..."
				}

				// Display the result
				fmt.Printf("Original file: %s\n", indexData.FileName)
				fmt.Printf("SimHash: %x\n", hash) // Print the SimHash of the matching chunk
				fmt.Printf("Byte offset: %d\n", offset)
				fmt.Printf("Phrase: %s\n", phrase)
				fmt.Println("----------")
			}
		}
	}
	return nil
}

func hammingdistance(a, b uint64) int {
	distance := 0
	for i := 0; i < 64; i++ {
		if (a & (1 << i)) != (b & (1 << i)) {
			distance++
		}
	}
	return distance
}
