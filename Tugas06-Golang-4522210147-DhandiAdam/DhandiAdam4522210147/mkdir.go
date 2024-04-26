package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("DhandiAdam", 1009)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'DhandiAdam'berhasil dibuat. ")
}
