package main

import "fmt"

// callback -> passar uma func como argumento(parametro)

func main() {
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55}...)
	fmt.Println("a soma dos numeros pares é", t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v % 2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}