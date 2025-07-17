package main

import (
	"fmt"
	"time"
)

/*

	Na pasta 5-range-close, lá a func main recebe o calor, enquanto a função que eu criei escreve

	worker() é uma função consumidora que lê valores do canal (RECEIVE).

	No main(), os valores são enviados para o canal (SEND).

	Regra importante:
	* Cada envio (SEND) bloqueia até que alguma goroutine faça a leitura (RECEIVE). Para evitar Deadlock

	Se houver somente 1 for-i (SEND - main) e somente 1 for-range (RECEIVE - go worker) o codigo fica limitado a 1 escrita 1 leitura, 
	
	o for-i bloqueia e espera até o for-range ler o canal

	O canal funciona como uma fila compartilhada:
	- A função main() envia valores.
	- Várias goroutines worker() leem esses valores ao mesmo tempo.

	Quando temos várias workers (go worker()):
	- Varias goroutines competem para ler o mesmo canal.
	- Assim, o canal é esvaziado mais rapidamente, permitindo processar tarefas "paralelamente"(Concorrente).

	1 Canal para varias Goroutines

	Esse padrão é chamado de Worker Pool Pattern e é muito usado em servidores web:
	- As requisições são enfileiradas (canal).
	- As workers (goroutines) retiram e processam as requisições de forma concorrente.

*/

func worker(workerId int, msg chan int) { // receive <-canal

	for res := range msg {
		fmt.Println("workerId>", workerId, "msg:",res)
		time.Sleep(time.Second)
	}

}

func main() {
	msg := make(chan int) // o canal é uma fila concorrente (thread-safe)

	// todo poder do go

	// ver a aula do wesley de novo
	// bota no 52 min

	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)

	// simboliza uma tarefa, onde varias Goroutines, go worker(), vão agilizar a entrega dela
	for i := 0; i < 10; i++ {
		msg <- i // send
	}
	// close(msg)
	
}