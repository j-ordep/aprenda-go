package main

import (
	"aprenda-go/cap26/ex01/cachorro"
	"fmt"
	"log"
)

/*
Crie um package "cachorro".
Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
Documente seu código com comentários, e utilize a função Idade na sua função main.
Rode seu programa para verificar se ele funciona.
Rode um local server com godoc e leia sua documentação.
*/

func main() {
	idadeDeCachorro, err := cachorro.Idade(5)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(idadeDeCachorro)
}