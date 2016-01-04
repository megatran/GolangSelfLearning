package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (z Square) area() float64 {
	return z.side * z.side
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	sq := Square{10}
	cr := Circle{5}
	info(sq)
	info(cr)
}
