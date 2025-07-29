package main

import "fmt"

func main() {
	resultado := Soma(10,10)
	fmt.Println(resultado)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

