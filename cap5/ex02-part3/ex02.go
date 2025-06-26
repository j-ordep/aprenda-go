package main

import "fmt"

// e se as variaveis tivessem tipos diferentes?
// Utilizo interface{} porque ela permite armazenar valores de qualquer tipo

/*
Declarar variável com tipo específico ->    var x Tipo = valor
Declarar e deixar o tipo ser inferido ->    x := valor
Converter um valor para outro tipo ->       x := Tipo(valor)
*/

func main() {

	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 100)
	d := (10 < 10)
	e := (10 >= 100)
	f := 10.5

	valores := []interface{}{a,b,c,d,e,f}

	for _, v := range valores {
		fmt.Printf("%v\n", v)
	}

}