package main

import "fmt"

/*Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let us know whether or not the argument was even. For example:
a. half(1) returns (0, false)
b. half(2) returns (1, true)
*/

func half(number int) (int, bool) {
	return number / 2, (number % 2) == 0
}

func main() {
	var n int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&n)
	fmt.Println(half(n))
}
