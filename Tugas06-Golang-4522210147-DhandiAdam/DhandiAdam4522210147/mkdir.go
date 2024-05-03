package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("DhandiAdam", 147)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'DhandiAdam'berhasil dibuat. ")
}
