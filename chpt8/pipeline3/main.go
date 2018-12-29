package main

import "sync"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
}

func printer(in <-chan int) {
	for x := range in {
		println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	var wg sync.WaitGroup

	go counter(naturals)
	wg.Add(3)
	go func() {
		defer wg.Done()
		squarer(squares, naturals)
	}()
	go func() {
		defer wg.Done()
		squarer(squares, naturals)
	}()
	go func() {
		wg.Wait()
		close(squares)
	}()
	printer(squares)
}
