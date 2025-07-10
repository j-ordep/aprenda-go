package main

import (
	"encoding/json"
	"fmt"
)

// Unmarshal

// converte JSON para código

type informacoes struct {
	Nome string `json:"Nome"`
	Idade int `json:"Idade"`
}

func main() {

	// Unarshal trabalha com []byte = "slice de byte"
	meuJson := []byte(`{"Nome": "James", "Idade":40}`)
	
	var pessoa1 informacoes

	// pessoa1 passa a ter os dados do JSON
	err := json.Unmarshal(meuJson, &pessoa1) // Unmarshal ( json_recebido -> variavel_struct )
	if err != nil {
		fmt.Println("erro:", err)
	}	

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.Idade)

}