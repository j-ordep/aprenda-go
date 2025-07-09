package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
}

type quadrado struct {
	base   float64
	altura float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.base * q.altura
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

func calcular(f figura) {
	fmt.Println(f)
	fmt.Println(f.area())
}

func main() {
	q := quadrado{4,4}
	c := circulo{2}

	calcular(q)
	calcular(c)

}