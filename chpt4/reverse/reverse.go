package main

import (
	"fmt"
)

func main() {
	data := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reverse(&data))
	fmt.Println(reverseBytes([]byte("abc")))
	fmt.Println(data)
}

func reverse(b *[5]int) []int {
	for i, j := 0, len(b)-1; i < j/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b[:]
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
