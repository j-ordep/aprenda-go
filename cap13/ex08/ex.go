package main

import "fmt"

// func que retorna uma func

func main() {

	a := retornafunc()

	a()

}

func retornafunc() func() {
	return func() {
		fmt.Println("Retornando uma func")
	}
}