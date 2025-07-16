package main

import (
	"fmt"
	"runtime"
	"sync"
)

// utiliza Mutex para resolver race condition do ex03-part1

var wb sync.WaitGroup
var mu sync.Mutex

func main() {
	contador := 0

	maxGoroutines := 100

	wb.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
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

	wb.Wait()

	fmt.Println("total:", contador)
}