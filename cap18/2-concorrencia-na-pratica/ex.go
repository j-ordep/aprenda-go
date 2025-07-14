package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Race Condition: 
		é quando duas ou mais goroutines acessam a mesma variável ao mesmo tempo, 
  		sem sincronização adequada, e pelo menos uma delas a modifica. Isso gera comportamento imprevisível.

	10 goroutines incrementam a variável "contador"

	yield -> altera para outra goroutine

	func main - goroutine
	go func() - goroutine

	"Ele pode rodar o i++ do for antes de rodar o v++ dentro da goroutine?"
	Sim, pode!
	Porque o for e as goroutines executam em contextos completamente diferentes e concorrentes.
*/

func main() {

	var wb sync.WaitGroup
	
	totalGoroutines := 1000

	contador := 0 // é o cara que as goroutine vão incrementar (RACE CONDITION)

	wb.Add(totalGoroutines) // vou usar esse numero de goroutines

	for i := 0; i < totalGoroutines; i++ {
		go func() { // goroutine

			v := contador

			// yield
			runtime.Gosched() // pausa e vai ver outra goroutine (provavelmente a main)
			
			v++
			contador = v
			wb.Done()
		}()
	}

	wb.Wait()
	fmt.Println(contador)
}