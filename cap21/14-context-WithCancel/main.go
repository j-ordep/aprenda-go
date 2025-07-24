package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// func WithCancel

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func ()  {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return

			default:
				n++
				time.Sleep(time.Microsecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())
	
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())
	cancel()
	fmt.Println("canceled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())
}