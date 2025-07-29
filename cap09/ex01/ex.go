package main

import "fmt"

func main() {

	num := [5]int {10, 20, 30, 40, 50}

	for i, v := range num {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", num)

}