package internals

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// RunFuzzy searches for nearly similar hashes in an index file and displays matching phrases from the original file.
//
// Parameters:
//   - indexFile: The path to the index file containing the precomputed SimHashes and their offsets.
//   - simHashStr: The SimHash value (in hexadecimal string format) to search for in the index.
//
// Returns:
//   - error: An error if any occurs during the execution, otherwise nil.
//
// The function performs the following steps:
//   1. Opens the index file and decodes its contents.
//   2. Checks if the original file specified in the index exists.
//   3. Parses the provided SimHash string into a uint64 value.
//   4. Opens the original file and reads chunks at the offsets where nearly similar hashes are found.
//   5. For each matching chunk, extracts and displays a phrase from the chunk along with the SimHash, byte offset, and original file name.
//   6. If no nearly similar hashes are found, it prints a message indicating so.
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

	// check instances if close similarity
	count := 0
	for hash, offsets := range indexData.Index {
		distance := hammingdistance(simHash, hash)
		if distance == 1 {
			count++
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

	if count == 0 {
		fmt.Println("No Nearly Similar Hashes found")
	}
	
	return nil
}

// The Hamming distance is the number of positions at which the corresponding bits are different.
// Parameters:
//   - a: the first 64-bit unsigned integer
//   - b: the second 64-bit unsigned integer
// Returns:
//   - int: the Hamming distance between the two integers
func hammingdistance(a, b uint64) int {
	distance := 0
	for i := 0; i < 64; i++ {
		if (a & (1 << i)) != (b & (1 << i)) {
			distance++
		}
	}
	return distance
}
