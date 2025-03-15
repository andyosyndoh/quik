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

// IndexData represents the structure for storing index information.
// It contains the following fields:
// - FileName: the name of the file being indexed.
// - ChunkSize: the size of each chunk in the file.
// - Index: a map where the key is a uint64 representing the chunk identifier,
//   and the value is a slice of int64 representing the positions within the chunk.
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

// NewFileIndex initializes a new FileIndex with specified chunk size and number of workers.
func NewFileIndex(chunkSize, numWorkers int) *FileIndex {
	return &FileIndex{
		chunkSize:  chunkSize,
		index:      NewIndex(),
		numWorkers: numWorkers,
	}
}