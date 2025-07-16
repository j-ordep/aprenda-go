package main

import "fmt"

func main() {

	// canal sem buffer
	
	canal1 := make(chan int)
	
	go func ()  {
		canal1 <- 10 // o canal sem buffer precisa ser passado em outra Goroutine, se não da Deadlock
	}()

	fmt.Println(<-canal1)
}