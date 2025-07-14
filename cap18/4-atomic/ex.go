package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// atomic

func main() {

	fmt.Println("Func1:")
	atomico1()
	fmt.Println("Func2:")
	atomico2()

}

func atomico1() {
	var contador int64
	contador = 0
	
	totalGoroutines := 1000

	var wb sync.WaitGroup
	wb.Add(totalGoroutines) 


	for i := 0; i < totalGoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wb.Done()
		}()
	}

	wb.Wait()
	fmt.Println(contador)
}

func atomico2() {

    var contador atomic.Uint64

    var wg sync.WaitGroup

    for range 50 {
        wg.Add(1)

        go func() {
            for range 1000 {

                contador.Add(1)
            }

            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", contador.Load())
}