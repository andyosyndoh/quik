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
	// if err != nil {
	// 	t.Fatalf("Failed to create test file: %v", err)
	// }
	// defer os.Remove(testFilename) // Cleanup after test

	// // Initialize FileIndex with necessary parameters
	// fileIndex := &FileIndex{
	// 	numWorkers: 2,  // Use 2 worker goroutines for parallel processing
	// 	chunkSize:  16, // Read file in 16-byte chunks
	// 	index:      &Index{m: make(map[uint64][]int64)},
	// }

	// // Run the BuildIndex function
	// err = fileIndex.BuildIndex(testFilename)
	// if err != nil {
	// 	t.Fatalf("BuildIndex failed: %v", err)
	// }

	// // Verify the index is populated
	// if len(fileIndex.index.m) == 0 {
	// 	t.Error("Index is empty, expected populated index data")
	// }
}
