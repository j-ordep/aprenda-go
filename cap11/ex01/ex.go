package main

import "fmt"

type pessoa struct {
	nome               string
	sobrenome          string
	sabores_de_sorvete []string
}

func main() {
	pessoa1 := pessoa{
		nome:               "Juliana",
		sobrenome:          "Martins",
		sabores_de_sorvete: []string{"menta, nutella"},
	}

	pessoa2 := pessoa{"João", "Pinheiro", []string{"Pistache", "doce de leite"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.sabores_de_sorvete {
		fmt.Println(v)
	}

	fmt.Println("")

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.sabores_de_sorvete {
		fmt.Println(v)
	}

}