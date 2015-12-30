package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E
	F
)
const (
	_ = iota      // throw away the 0
	G = iota * 10 // 1*10
	H = iota * 10 // 2 * 10
)

// bitwise operation
const (
	//bitwise shift operators. x << y means x × 2^y
	// while x >> y means x × 2^−y or, equivalently, x ÷ 2^y.

	_  = iota
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) //1 << (3*10)
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println("test bit shifting")
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
