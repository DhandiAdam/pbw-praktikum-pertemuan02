package main

import "fmt"

func main() {
	panjang := 5
	lebar := 3

	luas := hitungLuasPersegiPanjang(panjang, lebar)
	keliling := hitungKelilingPersegiPanjang(panjang, lebar)

	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
	fmt.Println("<<=============>>")
	fmt.Println("Penjelsan")
	fmt.Println("1. Membuat Function hitungLuasPersegiPanjang dengan paramater panjang,lebar int")
	fmt.Println("2. Mendeklarasi sebuah variabel bisa dengan var 'Nama Variabelnya apa' tipe datanya apa")
	fmt.Println("3. Atau bisa Langsung dengan Luas := objeknya apa")
	fmt.Println("4. Membuat variabel panjang := 5 dan Lebar := 3")
	fmt.Println("5. Lalu mendeklarasi variabel Luas := objek mengarah langsung ke function hitungLuasPersegiPanjang dan paramaternya(panjang, lebar) ")
	fmt.Println("6. Mendeklarasi variabel sama seperti no5 yaitu keliling := objek Mengarah langsung ke function yaitu keliling := hitungKelilingPersegiPanjang dengan paramater (panjang, lebar) ")
	fmt.Println("7. Lalu Function yang dituju me return/mendapatkan nilai dari hasil panjang x lebar")
}

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func hitungKelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
