// du1.go

/*
non concurrent
walks the directory tree and sums the sizes of all
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "." // or os.Args[1:]
	var nfiles int64
	var nbytes int64

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}
		if !info.IsDir() {
			nfiles++
			nbytes += info.Size()
		}
		return nil
	})

	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
