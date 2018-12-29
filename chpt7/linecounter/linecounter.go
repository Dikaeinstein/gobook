package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// LineCounter counts the lines in a byte slice.
type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	for input.Scan() {
		*l += LineCounter(1)
	}
	return int(*l), nil
}

func main() {
	var l LineCounter
	l.Write([]byte("hello"))
	fmt.Println(l) // "5", = len("hello")
	l = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&l, "hello, %s\n", name)
	fmt.Println(l) // "12", = len("hello, Dolly")
}
