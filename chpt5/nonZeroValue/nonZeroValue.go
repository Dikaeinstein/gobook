package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	nonZeroValue("https://golang.org")
}

// nonZero is a function that contains no return statement yet returns a non-zero value.
func nonZeroValue(url string) {
	type bailout struct{}
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	defer func() int {
		resp.Body.Close()
		switch p := recover(); p {
		case nil:
			fmt.Print(resp.Status)
			return 0
		case bailout{}:
			fmt.Print("bailed out of fetching resource")
			return 1
		default:
			panic(p) // Carry on panicking
		}
	}()
}
