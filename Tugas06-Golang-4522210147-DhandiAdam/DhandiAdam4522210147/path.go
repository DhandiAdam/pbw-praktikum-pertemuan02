package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileinfo, err := os.Stat("DhandiAdam")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if fileinfo.IsDir() {
		fmt.Println("File sudah ditemukan ")
	} else {
		fmt.Println("Sebuah File")
	}

}
