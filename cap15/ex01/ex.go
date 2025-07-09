package main

import "fmt"

// ponteiros

func main() {

	x := 10

	fmt.Println("valor de x:", x)
	fmt.Println("endereço de x:", &x)

	// y contem o endereço de x
	y := &x

	fmt.Println("valor de y:",*y) // imprime o valor que está no endereço y, ou seja o valor de x
	
}