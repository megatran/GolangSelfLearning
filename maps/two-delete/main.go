package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":  "Good morning",
		"one":   "Xin chao",
		"two":   "Bonjour",
		"three": "Buenos dias",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
