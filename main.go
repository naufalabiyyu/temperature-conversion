package main

import "fmt"

type celcius struct {
	suhu float64
}

type fahrenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

func (c celcius) toCelcius() float64 {
	return c.suhu
}

func (c celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32
}

func (c celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

func (f fahrenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

func (f fahrenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.suhu * (9.0 / 5.0)) - 459.67
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Pilih suhu awal : ")
	fmt.Println("1. Celcius ")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Masukan pilihan suhu awal : ")

	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)

	for suhuAwal < 1 || suhuAwal > 3 {
		fmt.Println("Suhu awal tidak tersedia, Pilihlah kembali suhu awal :")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("Pilih suhu akhir : ")
	fmt.Println("1. Celcius ")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Print("Masukan pilihan suhu akhir : ")

	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)

	for suhuAkhir < 1 || suhuAkhir > 3 {
		fmt.Println("Suhu akhir tidak tersedia, Pilihlah kembali suhu akhir :")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Print("Masukan suhu : ")
	fmt.Scanf("%f", &suhu)

	var interfaceSuhu hitungSuhu
	switch suhuAwal {
	case 1:
		interfaceSuhu = celcius{suhu}
	case 2:
		interfaceSuhu = fahrenheit{suhu}
	case 3:
		interfaceSuhu = kelvin{suhu}
	}

	var result float64

	switch suhuAkhir {
	case 1:
		result = interfaceSuhu.toCelcius()
	case 2:
		result = interfaceSuhu.toFahrenheit()
	case 3:
		result = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Hasil konversi : %.2f\n", result)

}
