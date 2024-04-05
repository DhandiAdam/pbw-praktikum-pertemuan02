package main

import (
	"fmt"
)

func main() {
	fmt.Println("Perkalian Factorial 7 ")
	fmt.Println(fact(7))

}

func fact(n int) int {
	if n == 0 || n == 1 {
		fmt.Printf("%d x 1\n ", n)
		return 1

	}
	fmt.Printf("%d X  ", n)
	result := n * fact(n-1)
	fmt.Printf("= %d\n", result)
	return result
}
