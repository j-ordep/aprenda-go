package main

import (
	"encoding/json"
	"fmt"
)

// Tem pegadinha!

// acho que ja sei!

/* 
	os atributos da struct estão com letra minuscula, 
	ou seja, são private.
	Precisam ser maiusculo para podermos ter acesso public
*/

type user struct {
	First string `json:"First"` // estava "first"
	Age   int `json:"Age"` // estava "age"
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println("slice de users:", users)

	j, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("json:", string(j))

}