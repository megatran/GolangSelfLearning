package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "inside scope"
		fmt.Println(y)
	}
	// fmt.Println(y) //outside scope so won't work
}
