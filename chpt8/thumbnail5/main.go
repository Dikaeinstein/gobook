package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

// makeThumbnails makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // Number of goroutines
	for f := range filenames {
		wg.Add(1)
		// Worker goroutine
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // Ignoring error
			sizes <- info.Size()
		}(f)
	}
	// closer goroutine
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	filenames := []string{
		"../Photos/IMG_0045.jpg",
		"../Photos/IMG_0046.jpg",
		"../Photos/IMG_0047.jpg",
		"../Photos/IMG_0048.jpg",
		"../Photos/IMG_0049.jpg",
		"../Photos/IMG_0050.jpg",
		"../Photos/IMG_0051.jpg",
	}
	ch := make(chan string)
	go func() {
		for _, f := range filenames {
			ch <- f
		}
		close(ch)
	}()
	bytesWritten := makeThumbnails(ch)
	fmt.Println(bytesWritten)
}
