package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dikaeinstein/gobook/chpt5/findlinks1"
)

var unseenLinks = make(chan string)

func main() {
	// Crawl the web breadth-first,
	worklist := make(chan []string)
	// starting from the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	breadthFirst(crawl, worklist)
}

func breadthFirst(f func(item string) []string, worklist chan []string) {
	seen := make(map[string]bool)
	for list := range worklist {
		for _, item := range list {
			if !seen[item] {
				seen[item] = true
				unseenLinks <- item
			}
		}
	}
	os.Exit(0)
}

func crawl(url string) []string {
	fmt.Println(url)
	links, err := findlinks1.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}
