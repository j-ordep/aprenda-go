package main

import "fmt"

// canais direcionais

func main() {
	c := make(chan int) // canal bidirecional
	cr := make(<-chan int) // canal receive
	cs := make(chan<- int) // canal send

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// posso transformal bidirecional em direciona
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// NÃO posso mudar um canal direciona de receive para send ou send para receive, muito menos para bidirecional
	// fmt.Printf("cs\t%T\n", (chan<- int)(cr)) // ERRO
}