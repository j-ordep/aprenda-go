package main

import "fmt"

// declarando variaveis (explicitamente tipadas) globais utilizando "var"
// atribuindo valor diretamente as variaveis
/* 
utilizando Sprintf e verbos de formatação 
	Sprintf pega a resposta formatada inteira (como se fosse Printf, mas NÃO IMPREME ELA),
	ele retorna uma única string, que é possivel guardar em uma variável e imprimir essa variavel para o resultado aparecer

	"%v" imprime o VALOR no formato padrão 
	"\t" separa na mesma linha como se utilizasse TAB
*/ 

var x int = 10
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y , z)
	fmt.Println(s)
}