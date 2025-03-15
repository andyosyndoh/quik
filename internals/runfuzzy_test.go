package internals

import (
	"encoding/gob"
	"os"
	"testing"
)

// TestRunFuzzy_ErrorCases tests different error scenarios in RunFuzzy
func TestRunFuzzy_ErrorCases(t *testing.T) {
	indexFile := "test_index.gob"
	simHash := "123abc" // Valid hex string

	// Test case: Index file does not exist
	err := RunFuzzy("nonexistent.gob", simHash)
	if err == nil {
		t.Errorf("Expected error for missing index file, got nil")
	}

	// Test case: Invalid SimHash format
	err = RunFuzzy(indexFile, "invalid-hash")
	if err == nil {
		t.Errorf("Expected error for invalid SimHash, got nil")
	}

	// Test case: Corrupted index file (invalid gob data)
	os.WriteFile(indexFile, []byte("corruptdata"), 0644)
	defer os.Remove(indexFile)
	err = RunFuzzy(indexFile, simHash)
	if err == nil {
		t.Errorf("Expected error for corrupted gob data, got nil")
	}

	// Test case: Original file missing
	createTestIndexFile(indexFile, "missing_file.txt", 16, map[uint64][]int64{0x123abc: {0}})
	err = RunFuzzy(indexFile, simHash)
	if err == nil {
		t.Errorf("Expected error for missing original file, got nil")
	}
}

// TestRunFuzzy_Success tests a case where "if distance == 1" is executed
func TestRunFuzzy_Success(t *testing.T) {
	indexFile := "test_index.gob"
	originalFile := "test_original.txt"
	simHash := "123abc"             // Test input hash
	hashInIndex := uint64(0x123abd) // Close to simHash (Hamming distance = 1)

	// Create original file
	content := "This is a test file for RunFuzzy."
	os.WriteFile(originalFile, []byte(content), 0644)
	defer os.Remove(originalFile)

	// Create test index file
	createTestIndexFile(indexFile, originalFile, 16, map[uint64][]int64{hashInIndex: {0}})
	defer os.Remove(indexFile)

	// Run test
	err := RunFuzzy(indexFile, simHash)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
}

// createTestIndexFile creates a valid gob-encoded index file
func createTestIndexFile(filename, originalFile string, chunkSize int, index map[uint64][]int64) {
	data := IndexData{
		FileName:  originalFile,
		ChunkSize: chunkSize,
		Index:     index,
	}

	file, _ := os.Create(filename)
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(data)
}
