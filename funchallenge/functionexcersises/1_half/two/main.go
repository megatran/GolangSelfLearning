package main

import "fmt"

func half(number int) (float64, bool) {
	return float64(number) / 2, number%2 == 0
}

func main() {
	var n int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&n)
	h, even := half(n)
	fmt.Println(h, even)
}
