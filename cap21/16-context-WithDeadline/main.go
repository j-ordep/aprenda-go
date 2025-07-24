package main

import (
	"context"
	"fmt"
	"time"
)

// Func WithDeadline

func main() {
	d := time.Now().Add(50 * time.Microsecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}