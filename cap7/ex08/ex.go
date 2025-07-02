package main

import "fmt"

func main() {

	dia := "sexta"

	switch {
	case dia == "segunda":
		fmt.Println("hoje é segunda") 
	case dia == "terça":
		fmt.Println("hoje é terça")
	case dia == "quarta":
		fmt.Println("hoje é quarta")
	case dia == "quinta":
		fmt.Println("hoje é quinta")
	case dia == "sexta":
		fmt.Println("hoje é sexta")
	}

}