package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	mirroredQuery()
}

func mirroredQuery() {
	responses := make(chan string, 3)
	ctx, cancelFunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(3)
	go func() { defer wg.Done(); responses <- request(ctx, "http://localhost:4000/?location=asia") }()
	go func() { defer wg.Done(); responses <- request(ctx, "http://localhost:4000/?location=europe") }()
	go func() { defer wg.Done(); responses <- request(ctx, "http://localhost:4000/?location=america") }()
	go func() {
		wg.Wait()
		close(responses)
	}()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancelFunc()
	}()
	for response := range responses {
		fmt.Println(response)
	}
	// return <-responses // return the quickest response
}

func request(ctx context.Context, hostname string) (response string) {
	done := make(chan struct{})
	errCh := make(chan error)
	var result []byte
	req, _ := http.NewRequest("GET", hostname, nil)
	req = req.WithContext(ctx)
	go func() {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()
		result, _ = ioutil.ReadAll(resp.Body)
		done <- struct{}{}
	}()
	for {
		select {
		case <-done:
			return string(result)
		case err := <-errCh:
			fmt.Println(err)
			return
		case <-ctx.Done():
			fmt.Println("Request cancelled!")
			return
		}
	}
}
