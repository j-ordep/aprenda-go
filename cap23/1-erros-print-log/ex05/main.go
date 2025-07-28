package main

import (
	"log"
	"os"
)

/*
	Panicln is equivalent to Println() followed by a call to panic().

	// http://godoc.org/builtin#panic
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}
}