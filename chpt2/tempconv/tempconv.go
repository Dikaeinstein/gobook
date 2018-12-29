// Package tempconv performs Celcius and Fahrenheit temperature computations
package tempconv

// Celsius unit of temperature
type Celsius float64

// Fahrenheit unit of temperature
type Fahrenheit float64

// Kelvin unit of temperature
type Kelvin float64

const (
	// AbsoluteC is the absolute temperature in Celsius
	AbsoluteC Celsius = -273.15
	// FreezingC is the freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC is the boiling temperature in Celsius
	BoilingC Celsius = 100
	// ZeroK is the Zero Kelvin temperature
	ZeroK Kelvin = -273.15
)
