package main

import "fmt"

func main() {

	si := []int{10, 20, 30, 40, 50}
	fmt.Println(soma(si...))
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}