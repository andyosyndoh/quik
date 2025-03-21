package internals

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// RunLookup performs a lookup operation on an index file using a provided SimHash string.
// It opens the index file, decodes the index data, verifies the existence of the original file,
// parses the SimHash string, and retrieves the byte offsets associated with the SimHash from the index.
// For each byte offset, it reads a chunk from the original file, extracts a phrase, and prints the
// original file name, byte offset, and the extracted phrase.
//
// Parameters:
//   - indexFile: The path to the index file.
//   - simHashStr: The SimHash string to lookup in the index.
//
// Returns:
//   - error: An error if any step of the lookup process fails, otherwise nil.
func RunLookup(indexFile, simHashStr string) error {
	// Open the index file for reading.
	dataFile, err := os.Open(indexFile)
	if err != nil {
		return fmt.Errorf("error opening index file: %v", err)
	}
	defer dataFile.Close()

	// Decode the index data from the file.
	var indexData IndexData
	decoder := gob.NewDecoder(dataFile)
	if err := decoder.Decode(&indexData); err != nil {
		return fmt.Errorf("error decoding index data: %v", err)
	}

	//Verify that the original file referenced in the index still exists.
	if _, err := os.Stat(indexData.FileName); os.IsNotExist(err) {
		return fmt.Errorf("original file %s not found", indexData.FileName)
	}

	//Parse the provided SimHash string into a uint64 value.
	simHash, err := strconv.ParseUint(simHashStr, 16, 64)
	if err != nil {
		return fmt.Errorf("invalid SimHash value: %v", err)
	}

	//Lookup the SimHash in the index to retrieve the byte offsets
	offsets, exists := indexData.Index[simHash]
	if !exists {
		return fmt.Errorf("SimHash not found in index: Ensure the file was indexed beforelooking up.")
	}

	file, err := os.Open(indexData.FileName)
	if err != nil {
		return fmt.Errorf("error opening original file: %v", err)
	}
	defer file.Close()

	for _, offset := range offsets {
		chunk := make([]byte, indexData.ChunkSize)
		n, err := file.ReadAt(chunk, offset)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading chunk at offset %d: %v", offset, err)
		}
		chunk = chunk[:n]

		// Convert chunk to string
		chunkStr := string(chunk)

		// Find first full word by skipping partial words
		startIdx := 0
		for startIdx < len(chunkStr) && (chunkStr[startIdx] != ' ' && chunkStr[startIdx] != '\n') {
			startIdx++
		}

		// Skip the space to reach the first full word
		if startIdx < len(chunkStr) {
			startIdx++
		}

		// Extract words from the valid start point
		words := strings.Fields(chunkStr[startIdx:])
		wordCount := len(words)

		// Determine length of phrase (at least full wordCount, up to 20 words)
		end := wordCount
		if end > 12 {
			end = 12
		}

		// Ensure we extract at least 20 full words in the chunk to build a phrase for the chunk
		phrase := strings.Join(words[:end], " ")

		if phrase == "" {
			end := min(len(chunkStr), 50)
			phrase = chunkStr[:end]
		}

		fmt.Printf("Original file: %s\n", indexData.FileName)
		fmt.Printf("Byte offset: %d\n", offset)
		fmt.Printf("Phrase: %s\n", phrase)
		fmt.Println("----------")

	}
	return nil
}
