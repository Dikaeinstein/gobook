package main

import (
	"fmt"

	"github.com/dikaeinstein/gobook/chpt2/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(2))
	fmt.Println(popcount.PCLoop(2))
}
