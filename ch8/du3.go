// du3.go
// concurrency limiting + waitgroup
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	root := "." // or from args
	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	wg.Add(1)
	go walkDir(root, &wg, fileSizes)

	// Close fileSizes after all walkDir goroutines complete
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var sema = make(chan struct{}, 20) // limit concurrency

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []fs.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}
