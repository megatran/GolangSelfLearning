package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		//%q is UTF-8 format
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
