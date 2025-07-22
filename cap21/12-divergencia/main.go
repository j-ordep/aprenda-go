package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
	
/* 2. Com throttling! Ou seja, com um número máximo de go funcs.

 	Dois canais.
	Uma func manda N números ao primeiro canal.
 	
	func outra():
	  - faz um for no canal1, criando X go funcs, cada uma com um range no primeiro canal que, para cada item, 
	    poe o retorno de trabalho() no canal dois.
	  - Como a func outra() é uma goroutine que cria várias goroutines internas, precisa usar WaitGroup para esperar todas elas terminarem
	    antes de fechar o canal e encerrar a goroutine externa.
    
	Trabalho() é um timer aleatório pra simular workload.
 	Por fim, range canal dois demonstra os valores.
*/

func main() {

	canal1 := make(chan int)
	canal2 := make(chan int)

	funcoes := 5 

	go manda(50, canal1)
	go outra(funcoes, canal1, canal2)

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

func outra(funcoes int, canal1, canal2 chan int) {

	var wb sync.WaitGroup 
	for i := 0; i < funcoes; i++ {
		wb.Add(1)
		go func ()  {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wb.Done()
		}()
	}

	wb.Wait() // espera terminar antes de fechar o canal
	close(canal2)
}

func trabalho(n int) int{
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}