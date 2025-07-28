package main

import (
	"fmt"
	"log"
)

// use o struct sqrt.Error como valor do tipo erro.

/* 
	A razão para criar uma struct como tipo de erro é adicionar mais informações ao erro — além da mensagem.

	Isso permite que você:
	  - Inclua dados extras (como localização, ID, usuário, status HTTP, etc).
	  - Preserve o erro original (err error) com detalhes técnicos reais.
	  - E ainda assim retorne um valor compatível com error, porque você implementa:
	  
*/

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		meuErro := fmt.Errorf("deu erro aqui doidão %v", f)
		return 0, sqrtError{"123", "321", meuErro}
	}
	return 42, nil
}