package main

import "fmt"

// closure -> captura um scope para que possamos utilizalo em outro contexto

func main() {

	a := i() // a = func() int { x++; return x }

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// meio que o scopo da variavel "x" esta na variavel "a", 
	// logo scopo de "b" o "x" tera um valor diferente

	b := i() // é um scopo diferente da variavel a, logo x tem um valor diferente
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0 // Variavel capturada pela CLOSURE. x esta fora do scopo da func abaixo, porem ainda guarda o valor
	return func() int {
		x++ // modifica a variavel capturada "x"
		return x
	}
}

/* Memory

----------------------------------------------

Clorure -> var a

x := 0 
	return func() int {
		x++ 
		return x
	}

---------------------------------------------- 

Valores diferentes para cada variavel, pois são diferentes scops

----------------------------------------------
Clorure -> var b

x := 0 
	return func() int {
		x++ 
		return x
	}

---------------------------------------------- 

*/