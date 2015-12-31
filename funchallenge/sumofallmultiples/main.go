package main

import "fmt"

func main() {
	limit := 0
	sum := 0
	fmt.Print("Please enter the max range of the natural numbers (i.e 1000): ")
	fmt.Scan(&limit)

	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of all multiples of 3 or 5 below ", limit, " is ", sum)
}
