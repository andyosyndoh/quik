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