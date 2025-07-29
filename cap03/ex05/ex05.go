package main

import "fmt"

// Conversions

// Criar variaveis "x" do tipo hotdog e "y" do tipo subjacente de hotdog, utilizando "var", no package-level scope (variaveis globais)
/*
	print valor da variavel "x"
	print tipo da variavel
	atribuir 42 a essa variavel "x" com operador "="
	print valor de x

	converter o tipo do valor de "x" para o tipo subjacente(int) e atribuir ao "y"
	print valor
	print tipo
*/ 

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Printf("Este é o valor inicial de x: %v\n", x)
	fmt.Printf("Este é o tipo de x: %T\n", x)
	x = 42
	fmt.Println("Atribuindo valor", x, "ao x")

	y = int(x) // converção do tipo do valor de "x" para int (subjacente de hotdog) e atribuindo ao "y"
	fmt.Println("Este é o valor de y:", y)
	fmt.Printf("Este é o tipo de y: %T", y) 
}