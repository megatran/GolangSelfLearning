package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zero", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
