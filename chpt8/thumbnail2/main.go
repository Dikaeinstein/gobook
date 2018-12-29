package main

import (
	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
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
	makeThumbnails(filenames)
}
