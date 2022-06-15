package main

import "fmt"

func main() {
	var umur float64
	fmt.Println("-Kategori Usia Anda-")
	fmt.Print("Masukan Umur Anda :")
	fmt.Scanln(&umur)

	fmt.Print("Umur anda masuk kedalam kategori :")
	if umur >= 0 && umur <= 2 {
		fmt.Println("Bayi")
	}
	if umur >= 3 && umur <= 5 {
		fmt.Println("Balita")
	}
	if umur >= 6 && umur <= 11 {
		fmt.Println("Bocil ff")
	}
	if umur >= 12 && umur <= 17 {
		fmt.Println("Bocah iykwim")
	}
	if umur >= 18 && umur <= 20 {
		fmt.Println("Fase berasa jadi Thomas Shelby")
	}
	if umur >= 21 && umur <= 24 {
		fmt.Println("Quarter life crisis a.k.a bocah edgy")
	}
	if umur >= 25 && umur <= 30 {
		fmt.Println("Budak Korporat :(")
	}
	if umur >= 31 && umur <= 41 {
		fmt.Println("Belum Nikah?,Hm?")
	}
	if umur == 42 {
		fmt.Println(";)")
	}
	if umur >= 43 && umur <= 50 {
		fmt.Println("Tentara apa yang paling kecil?-Tentara Sekutu. u ketawa? oh maaf pack")
	}
	if umur >= 51 && umur <= 68 {
		fmt.Println("Jamanku paling keren pokoe!!!1!")
	}
	if umur == 69 {
		fmt.Println("Kakek Legend")
	}
	if umur >= 70 && umur <= 100 {
		fmt.Println("Tobat Keck!")
	}
	if umur >= 101 {
		fmt.Println("prbbly ded.")
	}
	if umur <= -1 {
		fmt.Println("masi pejuh")
	}

}
