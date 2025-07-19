package main

import (
	"fmt"
)

// Chans par, ímpar, quit
// Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// Func receive é um select entre os três canais, encerra no quit

/* Problema!

	começa com 0 par e termina com 0 inpar

	RESOLVIDO em: /8-comma-ok

*/
func main() {
	par := make(chan int)
	inpar := make(chan int)
	quit := make(chan bool)

	go sendNum(par, inpar, quit)
	receiveNum(par, inpar, quit)

}

func sendNum(par, inpar chan int, quit chan bool) {
	for i := 0; i <= 30; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			inpar <- i
		}
	}
	close(par)
	close(inpar)
	quit <- true
}

func receiveNum(par, inpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println(v, "é par")

		case v := <-inpar:
			fmt.Println(v, "é inpar")

		case <-quit:
			return
		}
	}
} 

