// Package cachorro diz a idade do seu amigo em anos caninos
package cachorro

import (
	"errors"
)

// A função recebe a idade do cachorro, que deve ser do tipo int, e retorna a idade equivalente em anos caninos.
func Idade(i int) (int, error){

	idadeCanina := i * 7

	// Aqui verificamos se há algum erro na idade inserida
	if idadeCanina < 0 {
		return 0, errors.New("Idade inválida")
	}

	return idadeCanina, nil
}