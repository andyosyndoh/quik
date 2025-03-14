package internals

import (
	"hash/fnv"
	"io"
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
	wg.Add(fi.numWorkers)

	// Start worker goroutines to ensures efficient parallel processing of file chunks for SimHash computatio
	for i := 0; i < fi.numWorkers; i++ {
		go func() {
			defer wg.Done()
			h := fnv.New64a()
			for cd := range chunkChannel {
				simhash := computeSimHash(cd.data, h)
				resultChannel <- resultData{simhash, cd.offset}
			}
		}()

	}
	collectorDone := make(chan struct{})
	go func() {
		for rd := range resultChannel {
			fi.index.m[rd.simhash] = append(fi.index.m[rd.simhash], rd.offset)
		}
		close(collectorDone)
	}()

	offset := int64(0)
	buf := make([]byte, fi.chunkSize)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		data := make([]byte, n)
		copy(data, buf[:n])
		chunkChannel <- chunkData{data: data, offset: offset}
		offset += int64(n)
	}
	close(chunkChannel)
	wg.Wait()
	close(resultChannel)
	// <-collectorDone

	// return nil
}
