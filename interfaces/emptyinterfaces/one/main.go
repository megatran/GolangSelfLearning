package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	Maxspeed int
	Color    string
}

type car struct {
	vehicle
	wheel int
	Doors int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	mautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, bmw, boeing747, boeing757, boeing767, sanger, mautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
