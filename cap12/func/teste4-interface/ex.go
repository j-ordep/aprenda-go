package main

import (
	"fmt"
	"math"
)

// interfaces e polimorfismo

// aqui vemos que a função chama a interface, ou seja ela aceita as structs que implementam essa interface,
// no caso o retangulo(rect), mas se houver um circulo que tem como metodo as func de geometry 
// (ou seja, implementa a interface), essa func measure aceitaria o circulo tambem

// interface -> func (com receiver) -> struct

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


// a gente trabalha com a interface e não com a struct diretamente 
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{5, 2}
	c:= circle{5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}