package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/dikaeinstein/gobook/digest/sum/serial"
)

var p = flag.String("path", ".", "Root path to walk")

func main() {
	flag.Parse()
	mds, err := serial.MD5All(*p)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range mds {
		paths = append(paths, path)
	}
	sort.Strings(paths) // sort paths in place lexically
	for _, path := range paths {
		fmt.Printf("%x  %s\n", mds[path], path)
	}
}
