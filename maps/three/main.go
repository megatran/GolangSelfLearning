package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Xin chao",
		2: "Bonjour",
		3: "Buenos dias",
	}
	fmt.Println(myGreeting)

	//delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exsists")
		fmt.Println("val ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

}
