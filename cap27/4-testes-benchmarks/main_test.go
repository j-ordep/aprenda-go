package main

import "testing"

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(2, 2)
	}
}

func BenchmarkMultiplicacao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplicacao(2, 2)
	}
}