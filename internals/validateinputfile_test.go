package internals

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateInputFile(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() string // Setup function to create the test file
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid .txt file",
			setup: func() string {
				file := filepath.Join(t.TempDir(), "valid.txt")
				if err := os.WriteFile(file, []byte("valid content"), 0644); err != nil {
					t.Fatalf("Failed to create valid test file: %v", err)
				}
				return file
			},
			expectError: false,
		},
		{
			name: "File does not exist",
			setup: func() string {
				return filepath.Join(t.TempDir(), "nonexistent.txt")
			},
			expectError: true,
			errorMsg:    "input file does not exist",
		},
		{
			name: "Empty .txt file",
			setup: func() string {
				file := filepath.Join(t.TempDir(), "empty.txt")
				if err := os.WriteFile(file, []byte(""), 0644); err != nil {
					t.Fatalf("Failed to create empty test file: %v", err)
				}
				return file
			},
			expectError: true,
			errorMsg:    "input file is empty",
		},
		{
			name: "File is not a .txt file",
			setup: func() string {
				file := filepath.Join(t.TempDir(), "nontext.csv")
				if err := os.WriteFile(file, []byte("not a txt file"), 0644); err != nil {
					t.Fatalf("Failed to create non-txt test file: %v", err)
				}
				return file
			},
			expectError: true,
			errorMsg:    "input file must be a .txt file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := tt.setup()
			err := ValidateInputFile(file)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got: '%v'", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			}
		})
	}
}
