package main

import "fmt"

func main() {
	var myGreetings = make(map[string]string)
	myGreetings["Nhan"] = "Good morning"
	myGreetings["Tran"] = "Xin Chao"

	fmt.Println(myGreetings)
}
