package main

import "fmt"

// estilo poo

type pessoa struct {
	nome  string
	idade int
}

// receiver - a func mudeIdade() pertence a struct pessoa, como um metodo de uma class
func (p *pessoa) mudaIdade() {
	(*p).idade++
}

func mudaNome(novoNome string) *pessoa {
	p := &pessoa{nome: novoNome}
	return p
}

func main() {
	pessoa1 := pessoa{"Ju", 20}

	fmt.Println(pessoa1)

	pessoa1.mudaIdade()

	fmt.Println(pessoa1)

	novoNome := mudaNome("Juliana")

	fmt.Println(*novoNome)
}