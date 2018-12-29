// Package findlinks1 prints the links in a HTML document read from standard input
package findlinks1

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// func main() {
// 	doc, err := html.Parse(os.Stdin)
// 	// fmt.Println(doc)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
// 		os.Exit(1)
// 	}
// 	for _, link := range visit(nil, doc) {
// 		fmt.Println(link)
// 	}
// }

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link", "style":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img", "video":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

// func main() {
// 	for _, url := range os.Args[1:] {
// 		links, err := Extract(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
// 			continue
// 		}
// 		for _, link := range links {
// 			fmt.Println(link)
// 		}
// 	}
// }

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// CancelRequest is used by clients to terminate the HTTP request of Extract function
var CancelRequest context.CancelFunc
var Ctx context.Context

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	// resp, err := http.Get(url)
	Ctx, CancelRequest = context.WithCancel(context.Background())
	req, _ := http.NewRequest("GET", url, nil)
	req.WithContext(Ctx)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val) // Parse URL in the context of response request
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}
