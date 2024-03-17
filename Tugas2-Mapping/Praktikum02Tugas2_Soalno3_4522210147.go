package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	NPM     string
	Nama    string
	Jurusan string
}

func main() {
	// Daftar mahasiswa
	daftarMahasiswa := []Mahasiswa{
		{NPM: "4522210147", Nama: "Dhandi Adam", Jurusan: "Informatika"},
		{NPM: "4522210121", Nama: "Afiyah Nabila Putri", Jurusan: "Informatika"},
		{NPM: "4522210094", Nama: "Adam Kaze", Jurusan: "Informatika"},
	}

	cariMahasiswa(daftarMahasiswa)
}

func cariMahasiswa(daftar []Mahasiswa) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan NPM Mahasiswa yang ingin dicari: ")
	npm, _ := reader.ReadString('\n')
	npm = strings.TrimSpace(npm)

	found := false
	for _, mhs := range daftar {
		if mhs.NPM == npm {
			fmt.Println("Mahasiswa ditemukan:")
			fmt.Println("NPM    :", mhs.NPM)
			fmt.Println("Nama   :", mhs.Nama)
			fmt.Println("Jurusan:", mhs.Jurusan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Mahasiswa dengan NPM", npm, "tidak ditemukan.")
	}
}
