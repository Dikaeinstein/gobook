package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementMap: %v\n", err)
		os.Exit(1)
	}
	m := elementMap(map[string]int{}, doc)
	fmt.Println(m)
}

func elementMap(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = elementMap(m, c)
	}
	return m
}
