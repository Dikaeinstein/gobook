package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dikaeinstein/gobook/chpt2/tempconv"
)

func main() {
	for _, v := range os.Args[1:] {
		vfloat, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		f := tempconv.Fahrenheit(vfloat)
		c := tempconv.FToC(f)
		fmt.Printf("temp is %g°C\n", c)
	}
}
