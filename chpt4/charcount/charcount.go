// Package charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func charCount(in io.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of length of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	input := bufio.NewReader(in)
	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return counts, utflen, invalid
}
