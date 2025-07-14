
package main

import (
	"fmt"

	"sync"
)

// mutex resolve o problema de race condition

func main() {

	contador := 0 
	totalGoroutines := 1000

	var wb sync.WaitGroup
	wb.Add(totalGoroutines) 

	var mu sync.Mutex	

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			mu.Lock() 			// trava o acesso até 
			defer mu.Unlock() 	// libera acesso, quando tudo terminar (defer)
			v := contador
			v++ 				// seção crítica
			contador = v 
			wb.Done()
		}()
	}

	wb.Wait()
	fmt.Println(contador)
}