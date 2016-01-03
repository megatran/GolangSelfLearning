package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"Nhan", "Tran", 20}
	byteSlice, _ := json.Marshal(p1)
	fmt.Println("%T \n", byteSlice)
	fmt.Println(string(byteSlice))
}

//notExported data member is not in json since it's not uppercase
