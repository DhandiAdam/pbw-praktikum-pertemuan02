package main

import "fmt"

func main() {
	batasNilai := []int{90, 80, 70, 60}
	getNilaiHuruf := func(nilaiUjian int) string {
		for huruf, batas := range batasNilai {
			if nilaiUjian >= batas {
				switch huruf {
				case 0:
					return "A"
				case 1:
					return "B"
				case 2:
					return "C"
				case 3:
					return "D"
				}
			}
		}
		return "E"
	}

	// Menentukan nilai huruf untuk beberapa nilai ujian
	nilaiUjian := []int{85, 75, 95, 55, 65}
	for _, nilai := range nilaiUjian {
		fmt.Println("Nilai ujian", nilai, ":", getNilaiHuruf(nilai))
	}
}
