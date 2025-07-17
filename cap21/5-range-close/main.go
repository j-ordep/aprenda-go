package main

import (
	"fmt"
)

// chamar a função antes de rodar o for-range, pois o loop chama o canal, que precisa receber algum valor.
// logo a função tem que vir antes para dar tempo, mesmo sendo em Goroutines diferentes

// primeiro chamamos a função, canal<- enviamos o valor para o canal
// depois rodamos o for-range que retira valores do canal, <-canal


func main() {
	canal := make(chan int)
	
	go atribuiValorAoCanal(10, canal)

	for v := range canal { // receive <-canal
		fmt.Println("leu:", v)
	}

}

func atribuiValorAoCanal(j int, canal chan<- int) { // send
	for i := 0; i < j; i++ {
		canal <- i
		fmt.Println("escreveu:",i)
	}
	close(canal)
}