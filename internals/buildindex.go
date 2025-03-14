package internals

import (
	"hash"
	"strings"
)

type chunkData struct {
	data   []byte
	offset int64
}
type resultData struct {
	simhash uint64
	offset  int64
}

// computeSimHash generates a SimHash for the given data using a reusable hash.Hash64.
func computeSimHash(data []byte, h hash.Hash64) uint64 {
	text := string(data)
	words := strings.Fields(text)

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
}
