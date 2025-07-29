package main

import "fmt"

// loop for que mostra os valores ascii, decimal -> trasnforma em string, exibe hexadecimal e unicode tambem

func main() {

	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v \t%#x \t%U \n", x, string(x), x, x)
	}

}