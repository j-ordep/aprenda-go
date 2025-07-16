package main

import "fmt"

func main() {

	// canal com buffer

	canal2 := make(chan int, 2)

	canal2 <- 20
	canal2 <- 30

	fmt.Println(<-canal2) // 20
	fmt.Println(<-canal2) // 30
}