package internals

import (
	"os"
	"sync"
)

type chunkData struct {
	data   []byte
	offset int64
}
type resultData struct {
	simhash uint64
	offset  int64
}

// BuildIndex processes the file and builds the in-memory index.
func (fi *FileIndex) BuildIndex(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	chunkChannel := make(chan chunkData, 1000)
	resultChannel := make(chan resultData, 1000)

	var wg sync.WaitGroup
	// wg.Add(fi.numWorkers)
}
