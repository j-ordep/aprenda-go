package main

import "fmt"

/*
- Crie um struct "erroEspecial" que implemente a interface builtin.error. (tenha o metodo Error() string{})
- Crie uma função que tenha um valor do tipo error como parâmetro.
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/

type erroEspecial struct {
	qualquerCoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintln("Erro especial aqui!")
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte:", e.(erroEspecial).qualquerCoisa, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := erroEspecial{"Deu ruim maluco!"}
	erroComoParametro(arg)
}
