package main

import "fmt"

func main() {
	fmt.Println("Bubble Short")
	var arrayNumber [20]int
	data := [20]int{5, 3, 8, 1, 6, 9, 2, 7, 4} // Data yang akan dimasukkan

	// Memasukkan data ke dalam arrayNumber
	for i := 0; i < len(data); i++ {
		arrayNumber[i] = data[i]
	}

	// Menampilkan data sebelum pengurutan
	fmt.Println("Sebelum dilakukan pengurutan:")
	for _, val := range arrayNumber {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Algoritma Bubble Sort
	for i := 0; i < len(arrayNumber); {
		if i != (len(arrayNumber) - 1) {
			if arrayNumber[i] > arrayNumber[i+1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i+1]
				arrayNumber[i+1] = x
				i--
			}
		}
		if i > 0 {
			if arrayNumber[i] < arrayNumber[i-1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i-1]
				arrayNumber[i-1] = x
				i--
			}
		}
		i++
	}

	// Menampilkan data setelah pengurutan
	fmt.Println("Setelah dilakukan pengurutan:")
	for _, val := range arrayNumber {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
