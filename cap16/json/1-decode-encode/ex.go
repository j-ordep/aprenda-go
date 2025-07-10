package main

import (
	"encoding/json"
	"net/http"
)

// ENCODE E O DECODE 

type Pessoa struct {
	nome string
	idade int
}

// é como se fosse ResponseEntity e RequestBody
func criarPessoa(w http.ResponseWriter, r *http.Request) {

	var p Pessoa

	// pega o json do request.body e joga na variavel p
	err := json.NewDecoder(r.Body).Decode(&p) // como se fosse Unmarshal
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// lógica de processamento 
	// codigo codigo codigo 

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(p) // como se fosse marshal
}

func main() {
	http.HandleFunc("/pessoa", criarPessoa) // quando chamarem a rota "/pessoa" ele roda a func "criarPessoa"
	http.ListenAndServe(":8080", nil)
}