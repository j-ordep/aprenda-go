package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// WaitGroup é um contador de goroutines que permite esperar que todas terminem antes de continuar.
// Ele serve para sincronizar goroutines: você inicia várias tarefas e espera todas acabarem antes de seguir o programa.

// 1 goroutine por func

var wg sync.WaitGroup // precisa ser de escopo global

func main() {
	fmt.Println(runtime.NumCPU()) // mostra quantas threads meu pc tem

	fmt.Println("goroutine", runtime.NumGoroutine())

	// espera 2 tarefas (2 goroutines)
	wg.Add(2) // para cada valor add eu preciso de um done, no caso 2. *CHAME ANTES DAS GOROUTINES*
	
	go func1()
	go func2()

	// aqui temos 3 goroutines -> main - func1 - func2)
	fmt.Println("goroutine", runtime.NumGoroutine())

	wg.Wait() // bloqueia até o contador(add e done) chegar em zero, ou seja, espera até todas as tarefas terminarem

}

func func1() {
	for i := 0; i < 5; i++ {
		fmt.Println("func1", i)
		time.Sleep(40) // espera 40 ms
	}
	wg.Done() // diz ao contador: "terminei essa tarefa" - Internamente, é como fazer Add(-1)
}

func func2() {
	for i := 0; i < 5; i++ {
		fmt.Println("func2", i)
	}
	wg.Done()
}