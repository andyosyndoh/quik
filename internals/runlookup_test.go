package internals

import (
	"encoding/gob"
	"os"
	"testing"
)

// TestRunLookup tests the RunLookup function.
func TestRunLookup(t *testing.T) {
	// Create a temporary index file for testing.
	indexFile := "test_index.gob"
	originalFile := "test_original.txt"
	chunkSize := 10

	// Create a mock index data structure.
	indexData := IndexData{
		FileName:  originalFile,
		ChunkSize: chunkSize,
		Index: map[uint64][]int64{
			0x1234567890ABCDEF: {0, 20}, // SimHash and byte offsets
		},
	}

	// Write the mock index data to the index file.
	file, err := os.Create(indexFile)
	if err != nil {
		t.Fatalf("Failed to create index file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(indexData); err != nil {
		t.Fatalf("Failed to encode index data: %v", err)
	}

	// Create a mock original file with some content.
	originalContent := "This is a test file for RunLookup function. It contains some text for testing."
	err = os.WriteFile(originalFile, []byte(originalContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create original file: %v", err)
	}
	defer os.Remove(originalFile)

	// Define the SimHash to look up.
	simHashStr := "1234567890ABCDEF"

	// Call the RunLookup function.
	err = RunLookup(indexFile, simHashStr)
	if err != nil {
		t.Errorf("RunLookup failed: %v", err)
	}

	// Clean up the index file after the test.
	defer os.Remove(indexFile)
}