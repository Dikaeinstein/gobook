package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dikaeinstein/gobook/chpt5/foreachnode"
	"golang.org/x/net/html"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("text/html")
	if ct != "text/html" && strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return err
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	foreachnode.ForEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	// fetch page title
	fmt.Println(title(os.Args[1]))
}
