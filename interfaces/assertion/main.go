package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	//conversion
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	//assertion (interface only)
	fmt.Printf("%T\n", val.(int))
	fmt.Println(val.(int) + 6)
}
