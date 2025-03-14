package internals

import (
	"os"
	"testing"
)

// TestRunIndex verifies that RunIndex correctly processes an input file, builds an index, and serializes it.
func TestRunIndex(t *testing.T) {
	// Create a temporary input file
	inputFile := "test_input.txt"
	outputFile := "test_index.gob"
	content := "This is a test file for indexing. It contains multiple chunks of data."

	err := os.WriteFile(inputFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test input file: %v", err)
	}
	defer os.Remove(inputFile) // Cleanup after test
	defer os.Remove(outputFile)

	// // Define test cases
	tests := []struct {
		name       string
		chunkSize  int
		expectFail bool
	}{
		{"Valid chunk size", 16, false},
		{"Zero chunk size", 0, true},
		{"Negative chunk size", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunIndex(inputFile, tt.chunkSize, outputFile)
			if (err != nil) != tt.expectFail {
				t.Errorf("RunIndex() for %s failed: expected failure = %v, got error = %v", tt.name, tt.expectFail, err)
			}

			// Check if output file exists when expected
			if !tt.expectFail {
				if _, err := os.Stat(outputFile); os.IsNotExist(err) {
					t.Errorf("Expected output file %s was not created", outputFile)
				}
			}
		})
	}
}
