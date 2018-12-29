package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, \U00004e16\U0000754c"
	fmt.Println(len(s)) // "13" fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// "program" in Japanese katakana
	sj := "プログラム"
	fmt.Printf("% x\n", sj) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(sj)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))
}
