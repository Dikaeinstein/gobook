package main

import (
	"fmt"
)

func main() {
	const (
		KB = 1e3
		MB = 1e6
		GB = 1e9
		TB = 1e12
		PB = 1e15
		EB = 1e18
		ZB = 1e21
		YB = 1e24
	)

	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
