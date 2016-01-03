package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"Nhan", "Tran", 19}
	byteslice, _ := json.Marshal(p1)
	fmt.Println(string(byteslice))
}

/*
A tag for a field allows you to attach meta-information to the field which can be acquired using reflection. Usually it is used to provide transformation info on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.

The json package can look at the tags for the field and be told how to map json <=> struct field, and also extra options like whether it should ignore empty fields when serializing back to json.

Basically, any package can use reflection on the fields to look at tag values and act on those values. There is a little more info about them in the reflect package
http://golang.org/pkg/reflect/#StructTag :
 */