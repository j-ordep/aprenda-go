package main

import "fmt"

type pessoa struct {
	nome  string
  	idade int
}

type profissional struct {
	pessoa
	titulo string
	salario float64
}

func main() {
	pessoa1 := pessoa {
		nome: "João",
		idade: 20,
	}

	pessoa2 := profissional {
		pessoa: pessoa{
			nome: "Juliana",
			idade: 21,
		},
		titulo: "dentista",
		salario: 30.000,
	}

	pessoa3 := profissional {pessoa{"Karla", 45}, "Politica", 5.000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome, pessoa2.titulo)
	fmt.Println(pessoa3.nome, pessoa3.titulo)

}