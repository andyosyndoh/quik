package internals

import (
	"encoding/gob"
	"os"
	"testing"
)

// TestRunLookup tests the RunLookup function.
func TestRunLookup(t *testing.T) {
	// Define test file names
	indexFile := "test_index.gob"
	originalFile := "test_original.txt"
	chunkSize := 16

	// Create a mock original file
	originalContent := "This is a test file containing sample text for lookup."
	if err := os.WriteFile(originalFile, []byte(originalContent), 0644); err != nil {
		t.Fatalf("Failed to create original file: %v", err)
	}
	defer os.Remove(originalFile) // Clean up

	// Create a mock index data structure.
	indexData := IndexData{
		FileName:  originalFile,
		ChunkSize: chunkSize,
		Index: map[uint64][]int64{
			0xabcdef1234567890: {0}, // SimHash and byte offset
		},
	}

	// Write the mock index data to the index file.
	file, err := os.Create(indexFile)
	if err != nil {
		t.Fatalf("Failed to create index file: %v", err)
	}
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(indexData); err != nil {
		t.Fatalf("Failed to encode index data: %v", err)
	}
	file.Close()               // Close the file before testing
	defer os.Remove(indexFile) // Clean up

	// Define test cases
	tests := []struct {
		name      string
		indexFile string
		simHash   string
		wantErr   bool
	}{
		{
			name:      "Valid lookup",
			indexFile: indexFile,
			simHash:   "abcdef1234567890",
			wantErr:   false,
		},
		{
			name:      "Invalid SimHash",
			indexFile: indexFile,
			simHash:   "invalidhash",
			wantErr:   true,
		},
		{
			name:      "File not found",
			indexFile: "non_existent.gob",
			simHash:   "abcdef1234567890",
			wantErr:   true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunLookup(tt.indexFile, tt.simHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunLookup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
