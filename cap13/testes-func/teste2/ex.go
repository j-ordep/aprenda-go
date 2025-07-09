package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

// receiver
func (p *pessoa) informacoes() int {

	fmt.Println("Meu nome é", p.nome, ", eu tenho", p.idade, "anos")

	nascimento := 2025 - p.idade

	// espere
	defer fmt.Println("eu nasci no ano de", nascimento)

	return nascimento

}

func main() {

	pessoa1 := pessoa{"João", 20}

	// só consigo chamar o metodo informacoes se o tipo pertencer a struct pessoa 
	pessoa1.informacoes()


}