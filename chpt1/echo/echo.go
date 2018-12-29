// Echo prints its commandline arguments
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing line")
var sep = flag.String("s", " ", "separator")

func echo1() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func main() {
	// echo1()
	// echo2()
	echo4()
}
