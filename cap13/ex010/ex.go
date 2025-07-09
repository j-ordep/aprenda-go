package main

import "fmt"

func main() {

	a := closureFunc()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := closureFunc()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func closureFunc() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}