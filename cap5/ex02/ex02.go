package main

import "fmt"


func main() {

	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 100)
	d := (10 < 10)
	e := (10 >= 100)
	f := (10 > 10)

	valores := []bool{a,b,c,d,e,f}

	for _, v := range valores {
		fmt.Printf("%v\n", v)
	}

}