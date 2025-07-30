package main

import "fmt"

func main() {
	x := Soma(10,10)
	y := Multiplicacao(10,10)
	fmt.Println(x,y )
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func Multiplicacao(i ...int) int {
	total := 0
	for _, v := range i {
		total *= v
	}
	return total
}