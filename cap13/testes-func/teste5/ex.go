package main

import "fmt"



func main() {

	num := 100

	// função anonima
	func(x int) {
		fmt.Println("o numero é", num)
	}(num)

	fmt.Println("------")

	// função como expressão
	y := func(x int) int{
		dobro := x * 2
		fmt.Println("o numero é", num, "o dobro é", dobro)
		return dobro
	}

	y(num)


	// usando a func abaixo (func que retorna func)
	
	t := retornafunc() // func(i int) int {}
	h := t(3) // return 3 * 10

	fmt.Println(h)
}

// func que retorna func
func retornafunc() func(int) int {
	return func(i int) int {
		return i * 10
	}
}