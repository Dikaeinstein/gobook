package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

var format = flag.String("f", "jpg", "Output format")

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	flag.Parse()
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch *format {
	case "jpg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		err = gif.Encode(out, img, &gif.Options{})
	case "png":
		err = png.Encode(out, img)
	}
	return err
}
