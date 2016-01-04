package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"Nhan", "Last":"Tran", "Age":19}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}

/*
encode: take stuff from struct and send it out (writer interface)
decode: coming in and we're gonna read it (reader interface)

difference between marshal/unmarshal and encode/decode
- marshal/unmarshal: working with string or slice of bytes right inside application
- encode/decode deal with stuff outside (decode something comming in like stream or send it out to write it by encoding)
*/
