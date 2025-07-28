package main

import (
	"fmt"
	"log"
)

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go

type norgateMathError struct {
	lat string
	long string
	err error
}

// Implementa a interface error, retornando a mensagem completa
func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err) // loga o erro personalizado com localização + erro original (func nme)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// cria um erro interno com fmt.Errorf, incluindo o valor inválido
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f) // cria o erro com mensagem formatada

		// retorna erro customizado contendo o erro interno e localização
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}