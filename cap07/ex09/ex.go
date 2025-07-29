package main

import "fmt"

func main() {

	dia := "sexta"

	switch dia {
	case "segunda":
		fmt.Println("hoje é segunda") 
	case "terça":
		fmt.Println("hoje é terça")
	case "quarta":
		fmt.Println("hoje é quarta")
	case "quinta":
		fmt.Println("hoje é quinta")
	case "sexta":
		fmt.Println("hoje é sexta")
	}

}