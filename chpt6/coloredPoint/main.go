package main

import (
	"fmt"
	"image/color"

	"github.com/dikaeinstein/gobook/chpt6/geometry"
)

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geometry.Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{geometry.Point{X: 5, Y: 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

// ColoredPoint has a Point and Color
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
