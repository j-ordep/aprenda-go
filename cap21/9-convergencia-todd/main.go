package main

import (
	"fmt"
	"sync"
)

/* todd

Canais par, ímpar, e converge. 
Func send manda pares pra um, ímpares pro outro, depois fecha.
Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
Por fim um range retira todas as informações do canal converge.

*/

func main() {
	par := make(chan int)
	inpar := make(chan int)
	converge := make(chan int)

	go envia(par, inpar)
	go recebe(par, inpar, converge)

	for x := range converge {
		fmt.Println("Valor recebido:", x)
	}
}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {

	var wb sync.WaitGroup

	wb.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wb.Done()
	}()

	wb.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wb.Done()
	}()

	wb.Wait()
	close(c)
}