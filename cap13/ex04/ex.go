package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) mostrarInfo() {
	fmt.Println("Meu nome é", p.nome)
	fmt.Println("Minha idade é", p.idade)
}

func main() {
	pessoa1 := pessoa{"Juliana", 21}
	pessoa1.mostrarInfo()
}