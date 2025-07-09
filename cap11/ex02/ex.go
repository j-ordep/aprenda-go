package main

import "fmt"

type pessoa struct {
	nome               string
	sobrenome          string
	sabores_de_sorvete []string
}

func main() {
	meuMapa := make(map[string]pessoa)
	
	meuMapa["Martins"] = pessoa{
		nome:               "Juliana",
		sobrenome:          "Martins",
		sabores_de_sorvete: []string{"menta", "nutella"},
	}

	meuMapa["Pinheiro"] = pessoa{"João", "Pinheiro", []string{"Pistache", "doce de leite"}}

	for _, v := range meuMapa {
		fmt.Println("Meu nome é", v.nome, ", meus sorvetes favoritos são de:")
		for _, s := range v.sabores_de_sorvete {
			fmt.Println("-",s)
		}
	}

}