package main

import "fmt"

/*The optional <- operator specifies teh channel direction, send or receive. If no direction is given, the channel is bidrectional. A channel may be constrained only to send or only to received by converstion or assignment

chan T // can be used to to send or receive values of type T
chan<- float64 // can only be used to send float64s
<-chan int // can only be used to receive ints
*/

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	outtwo := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		outtwo <- sum
		close(outtwo)
	}()
	return outtwo
}
