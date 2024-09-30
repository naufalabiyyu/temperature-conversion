package main

func main() {

}

func CelciusToFahrenheit(suhuCelcius float64) float64 {
	suhuFahrenheit := (9 / 5 * suhuCelcius) + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	suhuKelvin := suhuCelcius + 273.15
	return suhuKelvin
}
