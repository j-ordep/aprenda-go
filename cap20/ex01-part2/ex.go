package main

import (
	"fmt"
	"sync"
)

// Alem da goroutine principal, crie duas outras goroutines.
// Cada goroutine adicional devem fazer um print separado.
// Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

// Jeito Complexo

var wb sync.WaitGroup

func main() {

	novaGoroutine(100)
	
	wb.Wait()
}

func novaGoroutine(i int) { 
	wb.Add(i) // i = 2

	for j := 0; j < i; j++ {
		x := j

		go func(i int) {
			fmt.Println("eu sou a goroutine numero:", i)
			wb.Done()
		}(x) // argumento = parametro, x = i
	}	
	
}