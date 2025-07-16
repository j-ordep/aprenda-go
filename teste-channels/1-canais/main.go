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

func numeros(canal chan<- int) { // chan <-  a func recebe um canal, no qual será atribuido um valor para este canal

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Microsecond * 150)
	}
	canal<- 10000 // canal recebe o valor
}

func letras(canal chan<- bool) { // chan <-  canal deve receber o valor 

	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Microsecond * 230)
	}
	canal<- true // canal recebe o valor
}

func main() {

	cn := make(chan int)
	cl := make(chan bool)

	go numeros(cn)
	go letras(cl)

	<-cn // pega o valor do canal
	<-cl // pega o valor do canal

	fmt.Println("Fim da execução!")
}