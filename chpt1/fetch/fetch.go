// Fetch prints the content found at a url
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// start := time.Now()
	// ch := make(chan string)

	// for _, url := range os.Args[1:] {
	// 	if !strings.HasPrefix(url, "http://") {
	// 		url = "http://" + url
	// 	}
	// 	go fetchall(url, ch) // start a goroutine
	// }
	// for range os.Args[1:] {
	// 	fmt.Fprintf(os.Stdout, "%s\n", <-ch) // receive from channel ch
	// }
	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fetch()
}

func fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		bytesWritten, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d\t status:%s\n", bytesWritten, resp.Status)
	}
}

func fetchall(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
