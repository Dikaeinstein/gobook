package main

import (
	"fmt"
	"io"

	"github.com/dikaeinstein/gobook/chpt7/bytecounter"
)

// NewWriter wraps the original writer, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer
type NewWriter struct {
	w            io.Writer
	bytesWritten int64
}

func (n *NewWriter) Write(p []byte) (int, error) {
	bytesWritten, err := n.w.Write(p)
	if err != nil {
		return 0, err
	}
	n.bytesWritten += int64(bytesWritten)
	return bytesWritten, nil
}

// CountingWriter given an io.Writer, returns a new Writer that wraps
// the original, and a pointer to an int64 variable that at any moment
// contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := NewWriter{w: w}
	return &newWriter, &newWriter.bytesWritten
}

func main() {
	var byteCounter bytecounter.ByteCounter
	b, c := CountingWriter(&byteCounter)
	fmt.Println(*c)
	b.Write([]byte("hello"))
	fmt.Println(*c)
}
