package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dikaeinstein/gobook/chpt5/findlinks1"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var n int // number of pending sends to worklist
var depth = flag.Int("depth", 3, "Limits the depth of the concurrent crawler")

func main() {
	flag.Parse()
	// Crawl the web breadth-first,
	worklist := make(chan []string)
	// starting from the command-line arguments.
	n++
	go func() { worklist <- os.Args[2:] }()
	breadthFirst(crawl, worklist)
}

func breadthFirst(f func(item string) []string, worklist chan []string) {
	counter := 0
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		if counter > *depth {
			break
		}
		list := <-worklist
		for _, item := range list {
			if !seen[item] {
				seen[item] = true
				n++
				go func(item string) { worklist <- f(item) }(item)
			}
		}
		counter++
	}
	os.Exit(0)
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	links, err := findlinks1.Extract(url)
	if err != nil {
		log.Print(err)
	}
	<-tokens // release the token
	return links
}
