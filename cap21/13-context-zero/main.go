package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:", ctx)
	fmt.Println("Context err:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
	fmt.Println("------------")

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("Context:", ctx)
	fmt.Println("Context err:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
}