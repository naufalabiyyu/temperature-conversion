package main

func main() {

}

func CelciusToFahrenheit(suhuCelcius float64) float64 {
	suhuFahrenheit := (9.0 / 5.0 * suhuCelcius) + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	suhuKelvin := suhuCelcius + 273.15
	return suhuKelvin
}

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	suhuCelcius := (suhuFahrenheit - 32) * (5.0 / 9.0)
	return suhuCelcius
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	suhuKelvin := (suhuFahrenheit + 459.67) * (5.0 / 9.0)
	return suhuKelvin
}

func KelvinToCelcius(suhuKelvin float64) float64 {
	suhuCelcius := suhuKelvin - 273.15
	return suhuCelcius
}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	suhuFahrenheit := (suhuKelvin * (9.0 / 5.0)) - 459.67
	return suhuFahrenheit
}
