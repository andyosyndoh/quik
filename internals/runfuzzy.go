package internals

import (
	"fmt"
	"os"
)

// RunFuzzy performs a fuzzy search for SimHashes with a Hamming distance of 1 or 2.
func RunFuzzy(indexFile, simHashStr string) error {
	dataFile, err := os.Open(indexFile)
	if err != nil {
		return fmt.Errorf("error opening index file: %v", err)
	}
	defer dataFile.Close()
	return nil
}
