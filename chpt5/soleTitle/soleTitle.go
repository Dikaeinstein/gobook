package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dikaeinstein/gobook/chpt5/foreachnode"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	title, err := soleTitle(doc)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	println(title)
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	foreachnode.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no tile element")
	}
	return title, nil
}
