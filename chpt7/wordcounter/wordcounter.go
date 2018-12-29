package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter counts words in a space separated byte slice
type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		*w += WordCounter(1)
	}
	return int(*w), nil
}

func main() {
	var w WordCounter
	w.Write([]byte("hello"))
	fmt.Println(w) // "5", = len("hello")
	w = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w) // "12", = len("hello, Dolly")
}
