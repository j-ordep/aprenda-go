package main

import "testing"

func TestSoma1(t *testing.T) {
	teste := Soma(10, 10)
	resultado := 20

	if teste != resultado {
		t.Error("expected:", resultado, "got:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := Soma(10, 11)
	resultado := 20

	if teste != resultado {
		t.Error("expected:", resultado, "got:", teste)
	}
}