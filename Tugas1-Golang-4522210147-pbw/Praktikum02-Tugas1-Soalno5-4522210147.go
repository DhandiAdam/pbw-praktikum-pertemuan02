package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	usiaStr := scanner.Text()

	usia, err := strconv.Atoi(usiaStr)
	if err != nil {
		fmt.Println("Usia harus berupa angka.")
		return
	}

	fmt.Println("OUTPUT")
	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Usia: %d\n", usia)
}
