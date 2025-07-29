package main

import "fmt"

// armazenando as variaveis em um []slice 
// percorrendo o slice com um "for-range" e imprimindo os valores

func main() {

	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 100)
	d := (10 < 10)
	e := (10 >= 100)
	f := (10 > 10)

	valores := []bool{a,b,c,d,e,f} // slice das variaveis

	for _, v := range valores {
		fmt.Printf("%v\n", v)
	}

}