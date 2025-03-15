package internals

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"
)

func TestRunFuzzy(t *testing.T) {
	tests := []struct {
		name       string
		indexData  IndexData
		simHashStr string
		expectErr  bool
	}{
		{
			name: "Successful fuzzy match",
			indexData: IndexData{
				FileName:  "testfile.txt",
				Index:     map[uint64][]int64{0b101010: {0}},
				ChunkSize: 50,
			},
			simHashStr: "2A", // 0b101010 in hexadecimal
			expectErr:  false,
		},
		{
			name: "SimHash not found",
			indexData: IndexData{
				FileName:  "testfile.txt",
				Index:     map[uint64][]int64{0b101010: {0}},
				ChunkSize: 50,
			},
			simHashStr: "FF",  // Unmatched SimHash
			expectErr:  false, // Should return "No Nearly Similar Hashes found"
		},
		{
			name: "Original file missing",
			indexData: IndexData{
				FileName:  "nonexistent.txt",
				Index:     map[uint64][]int64{0b101010: {0}},
				ChunkSize: 50,
			},
			simHashStr: "2A",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for the original file
			if tt.indexData.FileName != "nonexistent.txt" {
				if err := os.WriteFile(tt.indexData.FileName, []byte("This is a test file."), 0644); err != nil {
					t.Fatalf("Failed to create mock file: %v", err)
				}
				defer os.Remove(tt.indexData.FileName)
			}

			// Create the index file
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			if err := enc.Encode(tt.indexData); err != nil {
				t.Fatalf("Failed to encode index data: %v", err)
			}

			indexFileName := "indexfile.gob"
			if err := os.WriteFile(indexFileName, buf.Bytes(), 0644); err != nil {
				t.Fatalf("Failed to write index file: %v", err)
			}
			defer os.Remove(indexFileName)

			// Run the function
			err := RunFuzzy(indexFileName, tt.simHashStr)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
