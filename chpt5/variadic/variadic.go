package main

import (
	"fmt"
)

func sum(args ...int) int {
	total := 0
	if len(args) <= 0 {
		return 0
	}
	for _, arg := range args {
		total += arg
	}
	return total
}

func max(vals ...int) int {
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := 1
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func join(sep string, s ...string) string {
	start := ""
	for _, val := range s {
		start += (val + sep)
	}
	return start
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(join(" ", "Today", "is", "my", "birthday"))
}
