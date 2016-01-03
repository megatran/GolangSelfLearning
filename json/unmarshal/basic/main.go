package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person //initialize everything to zero
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	byteslice := []byte(`{"First":"Nhan", "Last":"Tran", "Age":19}`) //json
	json.Unmarshal(byteslice, &p1)

	fmt.Println("-----------------------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
