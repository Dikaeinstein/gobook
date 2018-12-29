package main

import "fmt"

func main() {
	fmt.Println(eliminateInplaceDup([]string{"a", "b", "b", "c", "d", "d", "d"}))
}

func eliminateInplaceDup(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			copy(strings[i:], strings[i+1:])
			strings = strings[:len(strings)-1]
		}
	}
	return strings[:len(strings)-1]
}
