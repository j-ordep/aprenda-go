package main

import "fmt"

func main() {
	canal := make(chan int)

	go send(canal)

	for v := range canal {
		fmt.Println(v)
	}
}

func send(canal chan int) {
	for i := 0; i < 100; i++ {
		canal <- i
	}
	close(canal)
}