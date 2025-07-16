package main

import (
	"fmt"
	"time"
)

/*	CICLO DO CHANNEL 

- Um envia outro recebe, enquanto um não se "move" o outro fica bloqueado (DeadLock)

	1. criar															=>		canal := make(chan bool)
	2. passar para uma outra Goroutine(func)							=> 		go outraFunc(canal)
	3. função recebe (chan<-) o canal como parametro					=> 		func outraFunc(canal chan<- bool) {}
	4. na outraFunc eu atribuo o valor ao canal							=>		canal<- true
	5. na main eu recebo o valor que foi passado na funcGoroutine		=>		<-canal

*/

// Buffered Channel

func numeros(canal chan<- int) {

	for i := 0; i < 10; i++ {
		canal <- i 			// 3 - envia valor para o canal
		fmt.Println("escrito no channel:", i)
		time.Sleep(time.Microsecond * 90)
	}
	fmt.Println("Fim da escrita!")
	close(canal) 			// 4 - Fecha o canal (encerra o fluxo)
}

func main() {

	cn := make(chan int, 3) 	// 1 - Cria canal (sem buffer)
	go numeros(cn)  		// 2 Goroutine escreve no canal

	for v := range cn { 	// 5 - Loop lê valores do canal - consome o <-canal (v <- cn)
		fmt.Println("lido do channel", v)
		time.Sleep(time.Microsecond * 180)
	}

	fmt.Println("Fim da execução!")
}