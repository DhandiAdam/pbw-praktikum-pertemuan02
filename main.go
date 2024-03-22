package main

import (
	"fmt"
	"progo/ratarata"
)

func main() {
	nilaiSiswa := []int{80, 75, 90, 85, 60}
	rataRata := ratarata.HitungRataRata(nilaiSiswa)
	fmt.Println("Rata Rata Nilai Mahasiswa: ", rataRata)
}
