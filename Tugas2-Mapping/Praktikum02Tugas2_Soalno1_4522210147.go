package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan Nama Anda : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukan Npm Anda: ")
	Npm, _ := reader.ReadString('\n')
	Npm = strings.TrimSpace(Npm)

	fmt.Print("Masukan Jurusan Anda ")
	Jurusan, _ := reader.ReadString('\n')
	Jurusan = strings.TrimSpace(Jurusan)

	BukuMap := map[string]string{
		"nama":    nama,
		"Npm":     Npm,
		"Jurusan": Jurusan,
	}
	fmt.Println(BukuMap)
	fmt.Println(BukuMap["nama"], nama)
	fmt.Println(BukuMap["Npm"], Npm)
	fmt.Println(BukuMap["Jurusan"], Jurusan)

}
