package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	nR := strings.NewReader("hello")
	lR := LimitReader(nR, 3)
	in := bufio.NewReader(lR)
	for i := 0; i < 3; i++ {
		b, _, err := in.ReadRune()
		if err == io.EOF && i == 2 {
			fmt.Println("We are at the end of the file")
		}
		fmt.Println(b)
	}
}

// LimReader is a Reader that wraps a Reader and reads up to limit
type LimReader struct {
	r     io.Reader
	limit int64
}

func (l *LimReader) Read(p []byte) (n int, err error) {
	input := bufio.NewScanner(l.r)
	input.Split(bufio.ScanRunes)
	for input.Scan() && int64(n) < l.limit {
		n++
	}
	return n, io.EOF
}

// LimitReader function in the io package accepts an io.Reader r
// and a number of bytes n, and returns another Reader that
// reads from r but reports an end-of-file conditionafternbytes
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimReader{r: r, limit: n}
}
