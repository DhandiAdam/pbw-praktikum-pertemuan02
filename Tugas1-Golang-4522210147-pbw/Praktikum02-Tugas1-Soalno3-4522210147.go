package main

import "fmt"

func main() {
	var nama string
	var usia int
	var kategori string
	fmt.Print("Masukan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukan usia : ")
	fmt.Scan(&usia)
	if usia < 18 {
		kategori = "Anak Anak"
		fmt.Println("Selamat Datang", nama, "Anda Termasuk Kategori", kategori)

	} else {
		kategori = "Dewasa"
		fmt.Println("Selamat Datang", nama, "Anda Termasuk Kategori", kategori)
	}
}
