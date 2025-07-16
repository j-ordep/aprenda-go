package main

import (
	"fmt"
	"runtime"
	"sync"
)

// utiliza Mutex para resolver race condition do ex03-part2

var wb sync.WaitGroup

var mu sync.Mutex

func main() {

	totalGoroutines := 100

	contador := criarGoroutine(totalGoroutines)

	wb.Wait()

	fmt.Println("total na main:", contador)

}

func criarGoroutine(numeroDeGoroutines int) int{

	wb.Add(numeroDeGoroutines)

	contador := 0

	for i := 0; i < numeroDeGoroutines; i++ {
		go func ()  {
			mu.Lock()
			defer mu.Unlock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println(contador)
			wb.Done()
		}()
	}

	fmt.Println("total na função de goroutines:", contador)

	return contador
}