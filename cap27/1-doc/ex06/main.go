
package main

import (
	"os"
)

// http://godoc.org/builtin#panic

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}
}

