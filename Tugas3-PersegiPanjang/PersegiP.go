package main

import "fmt"

func main() {
	panjang := 5
	lebar := 3

	luas, keliling := hitungLuasKelilingPersegiPanjang(panjang, lebar)

	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
	fmt.Println("<<<======================>>>")
	fmt.Println("Penjelasan")
	fmt.Println("1. Membuat Deklarasi Variabel itu ada berbagai cara")
	fmt.Println("2. Cara pertama dengan mendeklarasi sebuah variabel dengan var 'Namanya Apa' dan tipe datanya apa")
	fmt.Println("3. Cara kedua yaitu dengan cara panjang := objeknya apa")
	fmt.Println("Cara yang di nomer 3 hanya bisa didalam golang saja, kalau ada program yang bisa seperti itu selain golang mungkin ketidaktahuan Saya")
	fmt.Println("4. Mendeklarasi variabel panjang yaitu 5 dan lebar yaitu 3")
	fmt.Println("5. Membuat Function hitungLuasKelilingPersegiPanjang(panjang int, lebar int) (luas int, keliling int)")
	fmt.Println("6. dengan mendeklarasi variabel dalam sebuah paramater dan tidak perlu mendeklarasi variabel lagi dan langsung memasukan variabel untuk dihitung")
	fmt.Println("7. dikarenakan panjang itu 5 dan lebar itu 3 jadi kita masukan variabel rumusnya yaitu luas = panjang * lebar")
	fmt.Println("8. keliling = 2 * (panjang + lebar) dan langsung return untuk mendapatkan hasil yang ingin dimau")

}

func hitungLuasKelilingPersegiPanjang(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}
