package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* 1. Um stream vira centenas de go funcs que depois convergem.

 	Dois canais.
	Uma func manda N números ao primeiro canal.
 
	Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
	  - Como a func outra() é uma goroutine que cria várias goroutines internas, precisa usar WaitGroup para esperar todas elas terminarem
	    antes de fechar o canal e encerrar a goroutine externa.

	Trabalho() é um timer aleatório pra simular workload.
 	Por fim, range canal dois demonstra os valores.
*/

func main() {

	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
	
}

func manda(n int, canal1 chan int) {
	for i := 0; i < n; i++ {
		canal1 <- i
	}
	close(canal1)
}

func outra(canal1, canal2 chan int) {
	var wb sync.WaitGroup 

	for v := range canal1 {
		wb.Add(1)
		go func (x int) { // aqui eu gero 1 Goroutine para cada elemento do canal1 ()
			canal2 <- trabalho(x)
			wb.Done()
		}(v)
	}
	wb.Wait() // espera terminar antes de fechar o canal
	close(canal2)
}

func trabalho(n int) int{
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}