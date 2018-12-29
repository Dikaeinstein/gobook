package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dbServer := database{"shoe": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:5000", dbServer))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
