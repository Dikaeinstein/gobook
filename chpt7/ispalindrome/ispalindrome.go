package main

import (
	"fmt"
	"sort"
)

// IsPalindrome checks if s is a palindrome
func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i, j := 0, length-1; i < length/2; i++ {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
		j--
	}
	return true
}

type sortable []int

func (x sortable) Less(i, j int) bool { return x[i] < x[j] }
func (x sortable) Len() int           { return len(x) }
func (x sortable) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	intSlice := sortable{1, 2, 3, 4, 2, 1}
	fmt.Println(IsPalindrome(intSlice))
}
