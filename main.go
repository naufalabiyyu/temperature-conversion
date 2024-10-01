package main

import "fmt"

func main() {
	fmt.Println("Pilih kalkulator suhu berikut ini : ")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celcius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Print("Masukan pilihan anda : ")

	var pilihan int
	fmt.Scanf("%d", &pilihan)

	for pilihan < 1 || pilihan > 6 {
		fmt.Println("Kalkulator tidak tersedia, Pilih kembali kalkulator suhu :")
		fmt.Scanf("%d", &pilihan)
	}

	var suhu float64
	fmt.Print("Masukan suhu : ")
	fmt.Scanf("%f", &suhu)

	var result float64
	if pilihan == 1 {
		result = CelciusToFahrenheit(suhu)
	} else if pilihan == 2 {
		result = CelciusToKelvin(suhu)
	} else if pilihan == 3 {
		result = FahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		result = FahrenheitToKelvin(suhu)
	} else if pilihan == 5 {
		result = KelvinToCelcius(suhu)
	} else {
		result = KelvinToFahrenheit(suhu)
	}

	fmt.Printf("Hasil konversi : %.2f\n", result)

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
