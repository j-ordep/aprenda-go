package main

import "fmt"

// estilo procedural

type pessoa struct {
	nome  string
	idade int
}

func mudaIdade(p *pessoa) {
	(*p).idade++
}

func main() {
	pessoa1 := pessoa{"Ju", 20}

	fmt.Println(pessoa1)

	mudaIdade(&pessoa1)

	fmt.Println(pessoa1)
	
}