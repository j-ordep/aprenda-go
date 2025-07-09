package main

import "fmt"

// função que retorna uma função

func main() {

	x := retornafunc() // func(i int) int {}
	y := x(3)          // return 3 * 10

	fmt.Println(y)

	fmt.Println(retornafunc()(3)) // usar a func sem utilizar o x
}

// func que retorna func
func retornafunc() func(int) int {
	return func(i int) int {
		return i * 10
	}
}