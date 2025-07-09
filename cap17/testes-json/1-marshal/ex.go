package main

import (
	"encoding/json"
	"fmt"
)

// Marshal
// converte struct (ou map, slice ...) para json

type Pessoa struct {
	Nome string `json:"Nome"`
	Idade int `json:"Idade"`
}


func main() {
	
	pessoa := Pessoa{"James Bond", 40}

	p, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(p))
	fmt.Printf("tipo: %T\n", p)
	fmt.Println("se não transformar em string ele devolve isso:\n", p)
	

}