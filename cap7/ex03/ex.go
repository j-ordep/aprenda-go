package main

import "fmt"

// loop do tipo "while" que mostra do ano em que nasci até o ano atual

func main() {

	nascimento := 2005

	for nascimento <= 2025 {
		fmt.Println(nascimento)
		nascimento++
	}

}