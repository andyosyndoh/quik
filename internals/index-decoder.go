package internals

import (
	"encoding/gob"
	"fmt"
	"os"
)

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

	return nil
}
