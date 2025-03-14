package internals

import (
	"fmt"
	"os"
)

func IndexFileDecoder(file string) error {
	dataFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error opening index file: %w", err)
	}
	defer dataFile.Close()

	return nil
}
