package main

import "fmt"

//when the channel is closed, you can no longer put value on the channel, you can still receive values from the channel

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
