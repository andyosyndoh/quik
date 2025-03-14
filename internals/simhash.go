package internals

import (
	"hash"
	"strings"
)

// computeSimHash generates a SimHash for the given data using a reusable hash.Hash64.
func computeSimHash(data []byte, h hash.Hash64) uint64 {
	var simhash uint64
	var sums [64]int
	text := string(data)
	words := strings.Fields(text)
	return simhash
}
