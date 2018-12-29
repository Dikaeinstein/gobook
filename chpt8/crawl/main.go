package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dikaeinstein/gobook/chpt5/findlinks1"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()
	breadthFirst(crawl, worklist)
}

func breadthFirst(f func(item string) []string, worklist chan []string) {
	seen := make(map[string]bool)
	for list := range worklist {
		for _, item := range list {
			if !seen[item] {
				seen[item] = true
				go func(item string) { worklist <- f(item) }(item)
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
