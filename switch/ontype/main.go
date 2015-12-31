package main

import "fmt"

//switch on types
// --normall we switch on value of variable
// -- Go allows you to switch on type of variable

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { //asserting "is x of this type?"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
func main() {
	SwitchOnType(42)
	SwitchOnType("Tran")
	var test = Contact{"Nice to see you,", "Nhan"}
	SwitchOnType(test)
}
