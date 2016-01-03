package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Nhan", "Tran", 20, 007}
	byteSlice, _ := json.Marshal(p1)
	fmt.Println("%T \n", byteSlice)
	fmt.Println(string(byteSlice))
}

//notExported data member is not in json since it's not capitalized
