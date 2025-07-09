package main

import "fmt"

// função megablaster, parametro(variadico), multiplos retornos e sla mais o que

func main() {
	
	turno, total, quantidade := somar("tarde", 10, 10, 10)

	fmt.Println(turno, total, quantidade)
}

func somar(s string, x ...int) (string, int, int) { // "..." transgorma o parametro em slice, por isso posso usar range
	turno := ""

	if s == "manha" {
		turno = "Bom dia"
	} else if s == "tarde" {
		turno = "Boa tarde"
	} else {
		turno = "Boa noite"
	}

	somar := 0
	
	for _, v := range x {
		somar += v
	}

	return turno, somar, len(x)
}