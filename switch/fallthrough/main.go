package main

import "fmt"

func main() {
	name := "Nhan"
	switch name {
	case "Jack":
		fmt.Println("Wassup Jack")
	case "Nhan":
		fmt.Println("Wassup Nhan")
		fallthrough
	case "Tran":
		fmt.Println(" Your last name is Tran?")
	case "Sevy", "Natalie", "Mitch":
		fmt.Println("Wassup Robotics crew!")
	default:
		fmt.Println("You have no friends :(")
	}

	fmt.Println("Different test")
	lastname := "Tran"
	switch {
	case len(lastname) == 4:
		fmt.Println("Your lastname has 4 characters")
	case len(lastname) == 5:
		fmt.Println("Your lastname has 5 characters")
	default:
		fmt.Println("Your lastname doesn't match")
	}
}
