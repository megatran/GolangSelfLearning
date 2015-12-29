package main

import "fmt"

func main() {
	fmt.Println(42)
	//%d means base10 format
	//%b means base2 format
	//%x means base16 format
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
