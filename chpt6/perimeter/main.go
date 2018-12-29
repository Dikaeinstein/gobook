package main

import (
	"fmt"

	"github.com/dikaeinstein/gobook/chpt6/geometry"
)

func main() {
	perimeter := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}

	fmt.Println(perimeter.Distance()) // "12"
}
