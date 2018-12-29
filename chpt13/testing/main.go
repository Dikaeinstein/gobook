package main

import (
	"fmt"
	"math"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	x.a = true
	x.b = 4
	x.c = []int{1, 2, 3}
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println(unsafe.Alignof(x))
	fmt.Println(unsafe.Sizeof(x))
	fmt.Printf("%#016x\n", math.Float64bits(1.0))
}
