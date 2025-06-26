package main

import "fmt"

// criar um type, com tipo subjacente int
// cirar uma variavel (do tipo criado acima) com "var"
/*
	print valor da variavel "x"
	print tipo da variavel
	atribuir 42 a essa variavel "x" com operador "="
	print valor de x
*/ 

type hotdog int
var x hotdog

func main() {
	fmt.Printf("%v\n", x) // %v = valor
	fmt.Printf("O tipo de X é %T \n", x)
	x = 42
	fmt.Println(x)
}