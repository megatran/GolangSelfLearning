package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		//convert int32 to string to see the character, then take string and convert into slice of bytes. Can do that because string is a collection of runes
		//Each rune is a couple of bytes (1-4)
		fmt.Println(i, " - ", string(i), "-", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
