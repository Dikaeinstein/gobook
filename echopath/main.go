package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("p", 4000, "Port to listens on")

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on localhost:4000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	fmt.Fprintln(w, location)
}
