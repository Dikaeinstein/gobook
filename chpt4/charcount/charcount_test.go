package charcount

import (
	"fmt"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	input := strings.NewReader("test me\n")
	counts, utflen, invalid := charCount(input)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
