package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func ()  {
		c <- 10	
	}()
		
	v, ok := <-c
	fmt.Println(v, ok)		
	
	close(c)

	v, ok = <-c
	fmt.Println(v, ok) // 0 false (o canal foi já foi fechado)
}