package main

import (
	"log"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
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
