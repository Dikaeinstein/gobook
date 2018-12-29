package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("Error parsing html: %v", err)
	}
	elements := ElementsByTagName(doc, "head", "h1")
	println(elements[0].Data)
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	elements := []*html.Node{}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, name := range names {
			if c.Data == name {
				elements = append(elements, c)
				break
			}
		}
		elements = append(ElementsByTagName(c, names...), elements...)
	}
	return elements
}
