package internals

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
		return fmt.Errorf("SimHash not found in index")
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

		words := strings.Fields(string(chunk))
		phrase := strings.Join(words, " ")
		if len(phrase) > 50 {
			phrase = phrase[:50] + "..."
		}
	}
	return nil
}
