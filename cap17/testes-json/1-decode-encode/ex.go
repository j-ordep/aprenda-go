package main

import (
	"encoding/json"
	"net/http"
)

// testes json

// Assim que faz o ENCODE E O DECODE 

type Pessoa struct {
	nome string
	idade int
}

func criarPessoa(w http.ResponseWriter, r *http.Request) {

	var p Pessoa

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// lógica de processamento 
	// codigo codigo codigo 

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/pessoa", criarPessoa)
	http.ListenAndServe(":8080", nil)
}