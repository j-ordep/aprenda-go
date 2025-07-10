package main

import (
	"fmt"
	"sort"
)

// criando meu próprio sort
// complexo, podemos utilizar sort.Slice

type carro struct {
	nome string
	potencia int // cavalos
	consumo int // km/L
}

type ordenarPorPotencia []carro

func (p ordenarPorPotencia) Len() int {
	return len(p)
}

func (p ordenarPorPotencia) Less(i, j int) bool {
	return p[i].potencia < p[j].potencia
}

func (p ordenarPorPotencia) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}


type ordenarPorConsumo []carro

func (c ordenarPorConsumo) Len() int {
	return len(c)
}

func (c ordenarPorConsumo) Less(i, j int) bool {
	return c[i].consumo > c[j].consumo
}

func (c ordenarPorConsumo) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}



func main() {

	carros := []carro{
		{"La Ferrari", 600, 5},
		{"Fiat Uno", 75, 10},
		{"Honda Civic", 300, 2},
	}

	fmt.Println("inicial:\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Por potência\n",carros)
	
	sort.Sort(ordenarPorConsumo(carros))
	
	fmt.Println("Por consumo\n",carros)
}