package main

import (
	"fmt"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func dup(arr []int) int {
	start := 0
	for i, v := range arr {
		start ^= i ^ v
	}
	return start
}

func main() {
	fmt.Println(gcd(2, 3))
	fmt.Println(fibonacci(5))
	fmt.Println(dup([]int{1, 2, 3, 3, 4, 5, 6}))
}
