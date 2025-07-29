package main

import (
	"fmt"
	"log"
	"os"
)

/*

the Fatal function call os.Exit(1) after writing the log message

Fatalln is equivalent to Println() followed by a call to os.Exit(1)

*/

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("err happened", err)
		log.Fatalln(err)
		// panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}