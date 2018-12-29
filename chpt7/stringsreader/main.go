package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Reader implements the io.Reader interface by reading
// from a string.
type Reader struct {
	s        string
	i        int64
	prevRune int64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// NewReader Returns a new String reader that is parsable by html.Parse
func NewReader(s string) *Reader {
	return &Reader{s, 0, -1}
}

func main() {
	htmlString := "<!doctype html>\n<html><head><title>This is a test</title></head></html>"
	doc, err := html.Parse(NewReader(htmlString))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", doc)
}
