package main

import "fmt"

/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers
*/

//variadic parameters
func findgreatestnum(nums ...int) int {
	fmt.Printf("%T\n", nums)
	var max int
	for _, i := range nums {
		if max < i {
			max = i
		}
	}
	return max
}

func main() {
	greatest := findgreatestnum(1, 2, 3, 4)
	fmt.Print("Greatest number in the list is: ", greatest)
}
