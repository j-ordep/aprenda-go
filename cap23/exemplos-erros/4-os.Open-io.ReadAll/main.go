package main

import (
	"fmt"
	"io/ioutil"
	"os"
)




func main() {
	f, err := os.Open("names.txt") // abre o arquivo "names.txt"
	if err != nil {
		fmt.Println(err) // imprime o erro e encerra a execução com return
		return
	}
	defer f.Close()

	// ioutil foi desconsiderado (ainda funciona, mas está obsoleto)
	// recomendado é utilizar io.ReadAll()
	byteSlice, err := ioutil.ReadAll(f) // lê todo o conteúdo do arquivo retornando um slice de byte
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteSlice)) // imprime o conteúdo como string
}