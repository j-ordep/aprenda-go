package main

import (
	"fmt"
	"sync"
)

// Alem da goroutine principal, crie duas outras goroutines.
// Cada goroutine adicional devem fazer um print separado.
// Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

// Jeito simples

func main() {
	var wb sync.WaitGroup

	wb.Add(2)

	go func ()  {
		fmt.Println("Go func 1")
		wb.Done()
	}()

	go func ()  {
		fmt.Println("Go func 2")
		wb.Done()
	}()

	wb.Wait()
}
