package main

import "fmt"

func main() {
	result := fact(7)
	fmt.Printf("hasil dari perkalian factorial 7 adalah = %d\n", result)
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
