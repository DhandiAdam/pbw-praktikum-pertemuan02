package main

import "fmt"

func main() {
	var usia int
	var kategori string
	fmt.Print("Masukan usia : ")
	fmt.Scan(&usia)
	if usia < 18 {
		kategori = "Anak Anak"
		fmt.Println("Anda Termasuk Kategori : ", kategori)

	} else {
		fmt.Println("Usia Saya : ", usia)
		kategori = "Dewasa"
		fmt.Println("Kategori Usia : ", kategori)
	}
}
