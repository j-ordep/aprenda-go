package main

import "fmt"

// Crie um tipo para um struct chamado "pessoa"
// Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
// Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
// Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
// Demonstre no seu código:
// Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
// Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
// Se precisar de dicas, veja: https://gobyexample.com/interfaces


type humanos interface {
	falar()
}

type pessoa struct {
	nome string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("Olá meu nome é", p.nome, "minha idade é", p.idade)
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	pessoa1 := pessoa{"Juliana", 21}
	
	pessoa1.falar() // <- shortcut para (&pessoa1).falar()

	(&pessoa1).falar()

	dizerAlgumaCoisa(&pessoa1)
}