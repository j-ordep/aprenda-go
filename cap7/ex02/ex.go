package main

import "fmt"

// loop aninhado que tras a letras maiusculas do ascii de acordo com o numero

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}