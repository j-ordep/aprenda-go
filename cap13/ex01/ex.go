package main

import "fmt"

func main() {
	fmt.Println(soma(10,20))
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}