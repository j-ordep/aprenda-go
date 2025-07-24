package main

import (
	"fmt"
)

func main() {
	c := gen() // c só pode ser lido <-chan int 
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int { 
	c := make(chan int)

	go func () {
		for i := 0; i < 100; i++ {
			c <- i // SEND: canal <-
		}
		close(c)
	}()

	return c
}

func receive(canal <-chan int) {
	for v := range canal { // RECEIVE: <-canal
		fmt.Println(v)
	}
}