package main

import "fmt"

// callback (recebe func como argumento/parametro)

func main() {

	t := somentePares(soma, []int{10,20,30,40}...)
	fmt.Println(t)
	
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice[]int

	for _, v := range y {
		if v % 2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}