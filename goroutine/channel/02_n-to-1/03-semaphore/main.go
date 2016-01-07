package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true //put true onto channel
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true //put true onto channel
	}()

	go func() {
		<-done //throw value away off channel
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

/*
A semaphore is a way to lock a resource so that it is guaranteed that while a piece of code is executed, only this piece of code has access to that resource. This keeps two threads from concurrently accessing a resource, which can cause problems.
 */