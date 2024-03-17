package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PilihLanjut() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apakah Anda Ingin Mengakhiri Program (yes/no)?")
	pilihan, _ := reader.ReadString('\n')
	return strings.TrimSpace(pilihan)

}

func main() {
	for {
		fmt.Println("1. Hitung Luas Segitiga")
		fmt.Println("2. Hitung Persegi")
		fmt.Println("3. Program Berhenti")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Memilih : ")
		Memilih, _ := reader.ReadString('\n')
		Memilih = strings.TrimSpace(Memilih)

		switch Memilih {
		case "1":
			fmt.Println("Program Luas Segitiga")
			RumusLuasSegitiga()
		case "2":
			fmt.Println("Program Persegi")
			Persegi()
		case "3":
			fmt.Println("Program Berakhir")
			os.Exit(0)
		default:
			fmt.Println("Program Berakhir")

		}
		pilihan := PilihLanjut()
		if strings.ToLower(pilihan) == "no" {
			break
		}
	}

}

func HitungLuasSegitiga(inputAlas, TinggiAlas float64) float64 {
	return 0.5 * inputAlas * TinggiAlas
}

func RumusLuasSegitiga() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Alas Segitiga : ")
	inputAlas, _ := reader.ReadString('\n')
	inputAlas = strings.TrimSpace(inputAlas)

	LuasAlas, err := strconv.ParseFloat(inputAlas, 64)
	if err != nil {
		fmt.Print("Harus Angka Dan gk boleh Huruf ", err)
		return
	}

	fmt.Print("Masukan Tinggi Segitiga : ")
	TinggiAlas, _ := reader.ReadString('\n')
	TinggiAlas = strings.TrimSpace(TinggiAlas)

	LuasTinggi, err := strconv.ParseFloat(TinggiAlas, 64)
	if err != nil {
		fmt.Print("Harus Angka Dan gk boleh Huruf ", err)
		return
	}

	Luas := HitungLuasSegitiga(LuasAlas, LuasTinggi)
	fmt.Println("Luas Dari Segitiga ", Luas, "cm²")

}

func HitunglUASPersegi(SisiKiri, SisiKanan int) int {
	return SisiKiri * SisiKanan
}

func Persegi() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Sisi Kiri : ")
	InputSisiKiri, _ := reader.ReadString('\n')
	InputSisiKiri = strings.TrimSpace(InputSisiKiri)

	SisiKiri, err := strconv.Atoi(InputSisiKiri)
	if err != nil {
		fmt.Print("Harus Angka dan tidak boleh menggunakan Huruf ", err)
		return
	}

	fmt.Print("Masukan SisiKanan : ")
	InputSisiKanan, _ := reader.ReadString('\n')
	InputSisiKanan = strings.TrimSpace(InputSisiKanan)

	SisiKanan, err := strconv.Atoi(InputSisiKanan)

	if err != nil {
		fmt.Print("Harus Angka dan tidak boleh menggunakan Huruf ", err)
		return

	}

	LuasPersegi := HitunglUASPersegi(SisiKiri, SisiKanan)
	KelilingPersegi := 4 * SisiKiri

	fmt.Println("Luas Persegi Adalah : ", LuasPersegi, "cm²")
	fmt.Println("Keliling Persegi Adalah : ", KelilingPersegi, "cm")
}
