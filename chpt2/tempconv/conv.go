package tempconv

// CToF converts temperature from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts temperature from Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToF converts temperature from Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	c := KToC(k)
	return Fahrenheit(c*9/5 + 32)
}

// KToC converts temperature from Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k + ZeroK)
}
