package internals

import (
	"fmt"
	"os"
)


func RunLookup(indexFile, simHashStr string) error {
	// Open the index file for reading.
	dataFile, err := os.Open(indexFile)
	if err != nil {
		return fmt.Errorf("error opening index file: %v", err)
	}
	defer dataFile.Close()

	return nil
}
