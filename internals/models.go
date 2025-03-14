package internals

// Index holds the mapping from SimHash values to byte offsets.
type Index struct {
	m map[uint64][]int64
}

// FileIndex represents the indexing structure with configurable chunk size and workers.
type FileIndex struct {
	chunkSize  int
	index      *Index
	numWorkers int
}

type IndexData struct {
	FileName  string
	ChunkSize int
	Index     map[uint64][]int64
}

// NewIndex creates a new Index instance.
func NewIndex() *Index {
	return &Index{m: make(map[uint64][]int64)}
}

// Lookup retrieves the byte offsets for a given SimHash.
func (idx *Index) Lookup(hash uint64) []int64 {
	return idx.m[hash]
}