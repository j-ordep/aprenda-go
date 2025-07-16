package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// utiliza Atomic para resolver race condition do ex03-part2

var wb sync.WaitGroup
var contador int32

func main() {

	totalGoroutines := 100

	criarGoroutine(totalGoroutines)

	wb.Wait()

	fmt.Println("total na main:", contador)

}

func criarGoroutine(numeroDeGoroutines int) {

	wb.Add(numeroDeGoroutines)

	for i := 0; i < numeroDeGoroutines; i++ {
		go func ()  {
			atomic.AddInt32(&contador, 1)
			atomic.LoadInt32(&contador)
			fmt.Println(contador)
			wb.Done()
		}()
	}

	fmt.Println("total na função de goroutines:", contador)

}