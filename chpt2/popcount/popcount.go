package popcount

import "sync"

// pc[i] is the population count of i.
var pc [256]byte
var initializePcOnce sync.Once

func init() {
	initializePcOnce.Do(initializePc)
}

func initializePc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PCLoop returns the population count (number of set bits) of x.
func PCLoop(x uint64) int {
	var setBits int
	for i := 0; i < 8; i++ {
		setBits += int(pc[byte(x>>uint(i*8))])
	}
	return setBits
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	var setBits int
	for i := 0; i < 64; i++ {
		setBits += int((x >> uint(i)) & 1)
	}
	return setBits
}
