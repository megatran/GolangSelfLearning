package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 19
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 20
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//this doesn't work
	//fmt.Println(myAge + yourAge)

	//this conversion works
	fmt.Println(int(myAge) + yourAge)
}

/*
user defined types - we declare a new type foo
the underlying type of foo: int

conversion: int(myAge)
converting type foo to type int

This code is only for example. It's a bad practice to alias type
 */
