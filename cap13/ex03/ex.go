package main

import "fmt"

func main() {
	slaoq()
}

func slaoq() {
	defer fmt.Println("Esse texto tem defer, mas ele foi escrito em primeiro")
	fmt.Println("texto 2")
}