package main

import "fmt"

func main() {
	var suhu float64
	fmt.Print("Masukan suhu: ")
	fmt.Scanf("%f", &suhu)
	fmt.Printf("Suhu yg di input adalah, %f\n", suhu)

	var money uint64
	fmt.Print("Masukan uang: ")
	fmt.Scanf("%d", &money)
	fmt.Printf("Uang yg di input adalah, %d\n", money)

	var name string
	fmt.Print("Masukan nama: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Nama yg di input adalah, %s\n", name)
}
