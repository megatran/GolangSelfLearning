package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 //make the value of this address to  42
	fmt.Println(a)
}

//everything is passed by value in go
