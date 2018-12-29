package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "Show progress")

func main() {
	flag.Parse()
	roots := flag.Args()
	fileSizes := make(chan int64)
	dirFileSizes := make(chan int64)
	var wg sync.WaitGroup
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, root := range roots {
			// Traverse the file tree
			wg.Add(1)
			go walkDir(root, &wg, fileSizes, dirFileSizes)
		}
	}()
	var nBytes int64
	var nDirBytes int64
	var nFiles int64
	var tick <-chan time.Time
	// Print results periodically
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
		close(dirFileSizes)
	}()
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nBytes += size
			nFiles++
		case nDirBytes = <-dirFileSizes:
		case <-tick:
			if nDirBytes > 0 {
				println("Dir total")
				printDiskUsageMB(nFiles, nDirBytes)
			}
			printDiskUsage(nFiles, nBytes)
		}
	}

	// Print result
	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nFiles, float64(nBytes/1e9))
}

func printDiskUsageMB(nFiles, nBytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nFiles, float64(nBytes/1e3))
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64, dirFileSizes chan<- int64) {
	var nBytes int64
	defer n.Done()
	for _, entry := range dirent(dir) {
		if entry.IsDir() {
			subDir := path.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subDir, n, fileSizes, dirFileSizes)
		} else {
			size := entry.Size()
			fileSizes <- size
			nBytes += size
		}
	}
	dirFileSizes <- nBytes
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

func dirent(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
