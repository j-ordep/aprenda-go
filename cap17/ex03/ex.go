package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// minha parte

	uzischema := json.NewEncoder(os.Stdout) // Crie um Encoder que vai escrever a saída em formato JSON diretamente no terminal(os.Stdout).

	err := uzischema.Encode(users) // como uzischema é um encoder, ele tem acesso ao método Encode(), que transforma o parâmetro em JSON
	if err != nil {
		fmt.Println("ERRO:", err)
	}

	fmt.Println(users)

/* forma direta: 

	err := json.NewEncoder(os.Stdout).Encode(users) // method chaining
	if err != nil {
		fmt.Println("ERRO:", err)
	}

*/
	
}