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


// BuildIndex reads the specified file, processes it in chunks, and builds an index based on SimHash values.
// It uses multiple worker goroutines to compute SimHash values in parallel and a collector goroutine to
// aggregate the results into the index.
//
// Parameters:
//   filename: The path to the file to be indexed.
//
// Returns:
//   error: An error if any occurs during file reading or processing.
//
// The function performs the following steps:
// 1. Opens the specified file.
// 2. Creates channels for chunk data and result data.
// 3. Starts worker goroutines to process file chunks and compute SimHash values.
// 4. Starts a collector goroutine to aggregate the computed SimHash values into the index.
// 5. Reads the file in chunks and sends the chunks to the worker goroutines via the chunk channel.
// 6. Waits for all worker goroutines to finish processing.
// 7. Closes the result channel and waits for the collector goroutine to finish.
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
	<-collectorDone

	return nil
}
