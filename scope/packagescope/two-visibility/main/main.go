package main

import (
	"fmt"
	"github.com/megatran/GolangSelfLearning/scope/packagescope/two-visibility/vis"
)

func main() {
	x := 42
	fmt.Println(x)
	foo()
	fmt.Println(vistwo.MyName)
	vistwo.PrintVar()
}

func foo() {
	// can't access x
	//fmt.Println(x)
	fmt.Println(y)
}

//can still access outer scope even not in order
var y = 18
