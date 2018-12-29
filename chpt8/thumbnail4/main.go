package main

import (
	"gopl.io/ch8/thumbnail"
)

type item struct {
	thumbfile string
	err       error
}

// makeThumbnails makes thumbnails for the specified files in parallel.
// It returns the generated file names in an arbitrary order,
// or an error if any step failed.
func makeThumbnails(filenames []string) (thumbfiles []string, err error) {
	items := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			items <- it
		}(f)
	}
	for range filenames {
		it := <-items
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
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
