package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Hello",
		1: "Xin chao",
		2: "Bonjour",
		3: "Buenos dias",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
