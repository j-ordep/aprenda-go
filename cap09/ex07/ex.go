package main

import "fmt"

func main() {
	ss := [][]string{
		[]string{
			"Juliana",
			"Martins",
			"Odonto",
		},
		[]string{
			"João",
			"Pinheiro",
			"Dev",
		},
		[]string{
			"Karla",
			"Pinheiro",
			"Dormir",
		},
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}

}