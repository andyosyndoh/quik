package internals

import (
	"os"
	"testing"
)

// TestBuildIndex verifies that BuildIndex correctly processes a file and populates the index.
func TestBuildIndex(t *testing.T) {
	// Create a temporary test file
	testFilename := "testfile.txt"
	content := "This is a test file. It contains multiple words for indexing."
	err := os.WriteFile(testFilename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFilename) // Cleanup after test

	tests := []struct {
		name       string
		filename   string
		expectFail bool
	}{
		{"Valid file", testFilename, false},
		{"Non-existent file", "missingfile.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			// Initialize FileIndex with necessary parameters
			fileIndex := &FileIndex{
				numWorkers: 2,  // Use 2 worker goroutines for parallel processing
				chunkSize:  16, // Read file in 16-byte chunks
				index:      &Index{m: make(map[uint64][]int64)},
			}
		
			// Run the BuildIndex function
			err = fileIndex.BuildIndex(tt.filename)
			if (err != nil) != tt.expectFail {
				t.Fatalf("BuildIndex(%s) failed: expected failure = %v, got error = %v", tt.filename, tt.expectFail, err)
			}
		
			// Verify the index is populated
			if !tt.expectFail && len(fileIndex.index.m) == 0 {
				t.Error("Index is empty, expected populated index data")
			}
		})
	}

}
