package main

import "fmt"

type celsius struct {
	temperature float64
}

type fahrenheit struct {
	temperature float64
}

type kelvin struct {
	temperature float64
}

func (c celsius) toCelsius() float64 {
	return c.temperature
}

func (c celsius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.temperature) + 32
}

func (c celsius) toKelvin() float64 {
	return c.temperature + 273.15
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.temperature
}

func (f fahrenheit) toCelsius() float64 {
	return (f.temperature - 32) * (5.0 / 9.0)
}

func (f fahrenheit) toKelvin() float64 {
	return (f.temperature + 459.67) * (5.0 / 9.0)
}

func (k kelvin) toKelvin() float64 {
	return k.temperature
}

func (k kelvin) toCelsius() float64 {
	return k.temperature - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.temperature * (9.0 / 5.0)) - 459.67
}

type temperatureConverter interface {
	toCelsius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Choose initial temperature scale:")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Enter your choice for initial temperature scale: ")

	var initialScale int
	fmt.Scanf("%d", &initialScale)

	for initialScale < 1 || initialScale > 3 {
		fmt.Print("Invalid choice. Please choose the initial temperature scale again: ")
		fmt.Scanf("%d", &initialScale)
	}

	fmt.Println("Choose target temperature scale:")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Enter your choice for target temperature scale: ")

	var targetScale int
	fmt.Scanf("%d", &targetScale)

	for targetScale < 1 || targetScale > 3 {
		fmt.Print("Invalid choice. Please choose the target temperature scale again: ")
		fmt.Scanf("%d", &targetScale)
	}

	var temperature float64
	fmt.Print("Enter temperature value: ")
	fmt.Scanf("%f", &temperature)

	var converter temperatureConverter
	switch initialScale {
	case 1:
		converter = celsius{temperature}
	case 2:
		converter = fahrenheit{temperature}
	case 3:
		converter = kelvin{temperature}
	}

	var result float64
	switch targetScale {
	case 1:
		result = converter.toCelsius()
	case 2:
		result = converter.toFahrenheit()
	case 3:
		result = converter.toKelvin()
	}

	fmt.Printf("Conversion result: %.2f\n", result)
}
