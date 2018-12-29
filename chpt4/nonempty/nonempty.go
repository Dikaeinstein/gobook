package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(nonempty(data))
}

func nonempty(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}
