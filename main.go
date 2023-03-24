package main

import (
	"flag"
	"fmt"
	"github.com/NoraGith/funtemps2/tree/main/conv"
)

//definere variabler
var fahrenheit float64
var celsius float64
var kelvin float64
var out string

func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
}

//funksjon for flag
func main() {
	flag.Parse()
	if isFlagPassed("out") {
		if out == "C" && isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToCelsius(fahrenheit))
		} else if out == "F" && isFlagPassed("C") {
			fmt.Println(conv.CelsiusToFahrenheit(celsius))
		}
	}

	if isFlagPassed("out") {
		if out == "K" && isFlagPassed("C") {
			fmt.Println(conv.KelvinToCelsius(celsius))
		} else if out == "C" && isFlagPassed("K") {
			fmt.Println(conv.CelsiusToKelvin(kelvin))
		}
	}

	if isFlagPassed("out") {
		if out == "K" && isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToKelvin(kelvin))
		} else if out == "F" && isFlagPassed("K") {
			fmt.Println(conv.KelvinToFahrenheit(fahrenheit))
		}
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
