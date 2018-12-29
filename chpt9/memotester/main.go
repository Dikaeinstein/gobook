package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	memo "github.com/dikaeinstein/gobook/chpt9/memo1"
)

func httpGetBody(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func main() {
	incomingURLs := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}
	// Test sequential
	m := memo.New(httpGetBody)
	// for _, url := range incomingURLs {
	// 	start := time.Now()
	// 	value, err := m.Get(url)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Printf("%s, %s, %d bytes\n",
	// 		url, time.Since(start), len(value.([]byte)))
	// }
	// Test parallel
	m = memo.New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
