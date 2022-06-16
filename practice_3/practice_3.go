package main

import "fmt"

func main() {
	var angka, i int
	fmt.Print("Masukan Angka :")
	fmt.Scanln(&angka)
	fmt.Println("Faktor Bilangan", angka, "Adalah :")
	for i = angka; i >= 1; i-- {
		if angka%i == 0 {
			fmt.Println(i)
		}

	}
}
