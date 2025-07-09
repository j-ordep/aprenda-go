package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "joao",
		sobrenome: "pinheiro",
		fumante:   false,
	}

	cliente2 := cliente{ "juliana","martins",  false,}

	fmt.Println(cliente1, cliente2)

}