package main

import (
	"fmt"
	"os"
)

func main() {
	var err error

	err = os.Chmod("DhandiAdam", 0147)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("izin 'DhandiAdam' telah diubah menjadi 0147")
}
