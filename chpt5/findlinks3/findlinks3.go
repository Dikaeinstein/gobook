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
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
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
