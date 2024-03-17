package main

import "fmt"

func main() {

	var mahasiswa []string

	mahasiswa = append(mahasiswa, "Dhandi Adam")
	mahasiswa = append(mahasiswa, "Daffa Fathan")
	mahasiswa = append(mahasiswa, "Ozi")
	mahasiswa = append(mahasiswa, "Alpi")
	mahasiswa = append(mahasiswa, "Tegar")

	fmt.Println("Daftar Nama Mahasiswa:")
	for i, nama := range mahasiswa {
		fmt.Println("Perulang ke-", i+1, ":", nama)
	}
}
