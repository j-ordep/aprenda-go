package main

import "fmt"

func main() {

	x := 2005

	for {
		if x > 2025 {
			break
		}
		fmt.Println(x)
		x ++
	}

}