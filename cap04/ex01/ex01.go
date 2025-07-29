package main

import "fmt"

// tipo boleano e operadores relacionais
// == | <= | >= | != | < | >

var x bool

func main() {
	fmt.Println("Valor zero de x:", x) // zero value
	x =  true
	fmt.Println("Valor de x:", x)
	x = (10 < 100) // true
	y := (10 == 100) // false
	z := (10 != 100) // true 
	fmt.Printf("valor de y: %v\nvalor de z: %v", y, z)
}