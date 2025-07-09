package main

import "fmt"

// estilo procedural

type pessoa struct {
	nome  string
	idade int
}

// receiver - a func mudeMe() pertence a struct pessoa
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