package main

import "fmt"

type gente interface {
	oibomdia()
}

type pessoa struct {
	nome  string
	idade int
}

type dev struct {
	pessoa
	tecnologia string
	salario    float64
}

type dentista struct {
	pessoa
	area    string
	salario float64
}

// implementa "gente"
func (o dentista) oibomdia() {
	fmt.Println("Meu nome é", o.nome, ", que dente feio!")
}

// implementa "gente"
func (d dev) oibomdia() {
	fmt.Println("Meu nome é", d.nome, ", faz o L")
}

func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	dev := dev{pessoa{"João", 20},"go", 2000}
	serhumano(dev)

	dentista := dentista{pessoa{"Juliana", 21}, "arranca dente", 8000}
	serhumano(dentista)
}