package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt") // cria um arquivo "names.txt"
	if err != nil {
		fmt.Println(err) // imprime o erro e encerra a execução com return
		return
	}
	defer f.Close() // garante que o arquivo será fechado no final (nem roda)

	// cria um reader com a string "James"
	r := strings.NewReader("James") // transforma a string em um io.Reader

	io.Copy(f, r) // copia os dados de r para o arquivo f
}