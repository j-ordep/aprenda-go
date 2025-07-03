package main

import "fmt"

func main() {

	x := map[string][]string{

		"pinheiro_joao": []string{
			"codar", "go e java",
		},

		"martins_juliana": []string{
			"odonto", "joao",
		},
	}

	x["loureiro_kiko"] = []string {
		"Guitarra", "tacar fogo na guita",
	}

	for k, v := range x {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println(i, "-", h)
		}
	}

}