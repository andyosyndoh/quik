package internals

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ValidateInputFile checks if the provided input file meets certain criteria.
// It returns an error if any of the following conditions are met:
// - The file does not exist.
// - There is an error retrieving the file information.
// - The file is empty.
// - The file is not a .txt file.
//
// Parameters:
// - inputFile: The path to the input file to be validated.
//
// Returns:
// - An error if the input file does not meet the criteria, otherwise nil.
func ValidateInputFile(inputFile string) error {
	fileInfo, err := os.Stat(inputFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %v", err)
	}
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("input file is empty")
	}
	if strings.ToLower(filepath.Ext(inputFile)) != ".txt" {
		return fmt.Errorf("input file must be a .txt file")
	}
	return nil
}
