package main

import "fmt"

const (
	a = 2000 + iota
	b
	c
)

func main() {
	fmt.Println(a, b, c)
}