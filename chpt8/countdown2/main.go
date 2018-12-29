package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	// tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		// <-tick
		fmt.Println(countdown)
		select {
		case <-time.After(1 * time.Second):
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!!!")
}
