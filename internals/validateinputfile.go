package internals

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
