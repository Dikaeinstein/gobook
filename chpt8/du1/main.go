package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	fileSizes := make(chan int64)
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, root := range roots {
			// Traverse the file tree
			walkDir(root, fileSizes)
			close(fileSizes)
		}
	}()
	var nBytes int64
	var nFiles int64
	for fileSize := range fileSizes {
		nBytes += fileSize
		nFiles++
	}
	// Print result
	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirent(dir) {
		if entry.IsDir() {
			subDir := path.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirent(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
