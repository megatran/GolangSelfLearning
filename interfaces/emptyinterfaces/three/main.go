package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

//func specs takes any type (empty interface)
func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"Woof"}, true}
	fifi := cat{animal{"meow"}, true}
	shadow := dog{animal{"woof"}, true}
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)

}
