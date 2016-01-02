package main

//modify ex1 to use func expression

import "fmt"

func main() {
	var n int
	half := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}

	fmt.Print("Please enter a number: ")
	fmt.Scan(&n)
	h, even := half(n)
	fmt.Println(h, even)
}
