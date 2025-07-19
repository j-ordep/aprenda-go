package main 

import "fmt"

// Duas go funcs enviando X/2 numeros cada uma pra um canal
// For loop X valores, select case ←x

func main() {

	// cria 2 canais
	a := make(chan int)
	b := make(chan int)
	
	x := 100

	go func (x int)  {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func (x int)  {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}

}