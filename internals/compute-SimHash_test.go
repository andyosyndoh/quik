package internals

import (
	"hash/fnv"
	"testing"
)

// TestComputeSimHash checks the correctness of the computeSimHash function.
func TestComputeSimHash(t *testing.T) {
	h := fnv.New64() // Use FNV-1a hash for testing

	tests := []struct {
		input    string
		expected uint64 // expected SimHash value
	}{
		{"hello world", computeSimHash([]byte("hello world"), h)},
		{"hello hello world", computeSimHash([]byte("hello hello world"), h)},
		{"Go is awesome", computeSimHash([]byte("Go is awesome"), h)},
		{"", computeSimHash([]byte(""), h)}, // Edge case: empty string
		{"repeat repeat repeat", computeSimHash([]byte("repeat repeat repeat"), h)},
	}

	for _, test := range tests {
		h.Reset() // Reset hash before each test
		result := computeSimHash([]byte(test.input), h)

		if result != test.expected {
			t.Errorf("computeSimHash(%q) = %x; want %x", test.input, result, test.expected)
		}
	}
}
