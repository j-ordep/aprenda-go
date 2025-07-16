package main

import (
	"fmt"
	"runtime"
)

// Crie um programa que demonstra seu OS e ARCH.
// Rode-o com os seguintes comandos:
// go run
// go build
// go install

func main() {
	fmt.Println("seu OS é", runtime.GOOS)
	fmt.Println("seu ARCH é", runtime.GOARCH)
}