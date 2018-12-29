// Boiling prints the boiling point of water
package main

import (
	"fmt"

	"github.com/dikaeinstein/gobook/chpt2/tempconv"
)

const boilingF = tempconv.Fahrenheit(212.0)

func main() {
	f := boilingF
	c := tempconv.FToC(f)
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}
