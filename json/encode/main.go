package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Nhan", "Tran", 19, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}

/*
encode: take stuff from struct and send it out (writer interface)
decode: coming in and we're gonna read it (reader interface)

difference between marshal/unmarshal and encode/decode
- marshal/unmarshal: working with string or slice of bytes right inside application
- encode/decode deal with stuff outside (decode something comming in like stream or send it out to write it by encoding)
*/
