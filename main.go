package main

import (
	"flag"
	"fmt"
	"is105test/conv"
)

var fahrenheit float64
var celsius float64
var kelvin float64
var out string

func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "F", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "F", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
}

func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°C", fahrenheit, conv.FahrenheitToCelsius(fahrenheit))
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°K", fahrenheit, conv.FahrenheitToKelvin(fahrenheit))
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Printf("%v°F er = %.2f°C", celsius, conv.CelsiusToKelvin(celsius))
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Printf("%v°F er = %.2f°C", celsius, conv.CelsiusToFahrenheit(celsius))
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %.2f°C", kelvin, conv.KelvinToCelsius(kelvin))
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %.2f°F", kelvin, conv.KelvinToFahrenheit(kelvin))
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
