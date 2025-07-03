package main

import "fmt"

func main() {

	amigos := map[string]int{
		"juliana": 989301245,
		"joao":    998686167,
	}

	fmt.Println(amigos)

	fmt.Println(amigos["juliana"])

	amigos["gopher"] = 4444

	fmt.Println(amigos)

}