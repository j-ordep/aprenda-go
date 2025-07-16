package main

import "fmt"

// canais direcionais
// transformando um canal bidirecional em canais send ou receive nas funções

func main() {
	canal := make(chan int) // canal bidirecional

	go send(canal) // pelo menos uma das funções deve rodar em outra goroutine	
	receive(canal)

}

func send(s chan<- int) { // somente envia para o canal
	s <- 10
}

func receive(r <-chan int) { // somente retira do canal
	fmt.Println("O valor recebido do canal foi:",<- r)

}