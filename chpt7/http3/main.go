package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoe": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.Handle("/", http.HandlerFunc(db.notFound))
	log.Fatal(http.ListenAndServe("localhost:5000", mux))
}

var rwMu sync.RWMutex

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	var itemList = template.Must(template.New("itemlist").Parse(
		`<h3>Item List</h3>
		<table>
			<tr style='text-align: left'>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range $key, $value := .}}
			<tr>
				<td>{{ $key }}</td>
				<td>{{ $value }}</td>
			</tr>
		</table>
		{{end}}`,
	))

	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil {
		msg := fmt.Sprintf("Price: %f is invalid\n", price)
		http.Error(w, msg, http.StatusBadRequest)
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "Successfully created item: %s\n", item)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil {
		msg := fmt.Sprintf("Price: %f is invalid\n", price)
		http.Error(w, msg, http.StatusBadRequest)
	}
	rwMu.RLock()
	_, ok := db[item]
	rwMu.RUnlock()

	if !ok {
		msg := fmt.Sprintln("Item not found")
		http.Error(w, msg, http.StatusNotFound)
	}
	rwMu.Lock()
	defer rwMu.Unlock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "Successfully updated item: %s\n", item)
}

func (db database) notFound(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("no such page: %s\n", r.URL)
	http.Error(w, msg, http.StatusNotFound)
}
