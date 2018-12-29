package main

import (
	"github.com/dikaeinstein/gobook/chpt2/popcount"
)

func countDiffBits(s1, s2 [32]byte) int {
	count := 0
	for i, value := range s1 {
		pc1 := popcount.PopCount(uint64(value))
		pc2 := popcount.PopCount(uint64(s2[i]))

		if pc1 != pc2 {
			count++
		}
	}
	return count
}
