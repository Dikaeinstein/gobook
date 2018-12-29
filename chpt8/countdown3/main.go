package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	countdown := 10
	for countdown > 0 {
		select {
		case <-tick:
			fmt.Println(countdown)
			countdown--
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
			// fmt.Println("ticking!")
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!!!")
}
