package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f")
}

func f() {
	defer func ()  {
		if r := recover(); r != nil { // retorna para a func main
			fmt.Println("Recovered in f")
		}
	}()

	fmt.Println("Calling g")
	g(0)

	fmt.Println("Returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("defer in g", i)

	fmt.Println("Printing in g", i)

	g(i + 1)
}