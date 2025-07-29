package main

import "fmt"

// Analisando a string "s"
// passando um for que percorre cada caracter ( RUNE ) da string
// retornando o valor, tipo, valor unicode e valor hexadecimal

// rune é semelhante a char. caracteristicas do rune -> 32 bits - unicode(utf8) - int32

func main() {

	s := "Hello playground"

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

}