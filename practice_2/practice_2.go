package main

import "fmt"

func main() {
	var nama string
	var nilai float64 //antisipasi guru/dosen memberi nilai(koma).
	fmt.Println("-Nilai Ujian-")

	fmt.Print("Masukan Nama Mahasiswa :")
	fmt.Scanln(&nama)

	fmt.Print("Masukan Nilai Mahasiswa :")
	fmt.Scanln(&nilai)

	fmt.Println("Nama Mahasiswa :", nama)

	if nilai >= 0 && nilai <= 34 {
		fmt.Println("Nilai : D")
	}
	if nilai >= 35 && nilai <= 49 {
		fmt.Println("Nilai : C")
	}
	if nilai >= 50 && nilai <= 64 {
		fmt.Println("Nilai : B")
	}
	if nilai >= 65 && nilai <= 79 {
		fmt.Println("Nilai : B+")
	}
	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Nilai : A")
	}
}
