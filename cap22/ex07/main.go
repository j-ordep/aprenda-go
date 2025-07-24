package main

import (
	"fmt"
	"sync"
)

func main() {
	canal := make(chan int)

	go minhaFunc(canal)

	for v := range canal {
		fmt.Println(v)
	}
}

func minhaFunc(canal chan int) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // adiciona uma goroutine ao contador
		go func() {
			defer wg.Done() // marca como concluída quando terminar
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}

	wg.Wait()
	close(canal)
}
