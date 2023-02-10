package conv


// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius := (value - 32)*(5/9)
	return Celsius
}

func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit := value *(9/5) + 32
	return Fahrenheit
}

func  KelvinToFahrenheit(value float64) float64 {
  Fahrenheit := (value - 273.15)*(9/5) + 32
	return Fahrenheit
}

