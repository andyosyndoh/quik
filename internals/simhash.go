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

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	for word, cnt := range counts {
		h.Reset()
		h.Write([]byte(word))
		hash := h.Sum64()

		for i := 0; i < 64; i++ {
			if (hash & (1 << i)) != 0 {
				sums[i] += cnt
			} else {
				sums[i] -= cnt
			}
		}
	}

	for i := 0; i < 64; i++ {
		if sums[i] >= 0 {
			simhash |= 1 << i
		}
	}

	return simhash
}
