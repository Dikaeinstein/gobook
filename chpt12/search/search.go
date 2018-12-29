package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dikaeinstein/gobook/chpt12/params"
)

func main() {
	http.Handle("/search", http.HandlerFunc(search))
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}
	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
