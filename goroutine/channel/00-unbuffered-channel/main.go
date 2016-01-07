package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //unbuffered channel

	go func() {
		for i := 0; i < 10; i++ {
			c <- i //this stops until the func below takes the value
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) //this stops until the upper func passes the value
		}
	}()
	time.Sleep(time.Second)
}

/*
The <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bi-directional. A channel may be constrained only to send or only to receive by conversion or assignment.
*/
