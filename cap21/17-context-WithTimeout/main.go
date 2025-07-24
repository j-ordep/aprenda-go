package main

import (
	"context"
	"fmt"
	"time"
)

// Func WithTimeout

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Microsecond)

	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}