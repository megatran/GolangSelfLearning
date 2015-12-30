package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) //address in func zero
	*z = 0
}
func main() {
	x := 5
	fmt.Printf("%p\n", &x) //address in main
	fmt.Println(&x) // address in main
	zero(&x)
	fmt.Println(x)
}
