package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wb sync.WaitGroup

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