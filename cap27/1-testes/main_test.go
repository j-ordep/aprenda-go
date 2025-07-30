package main

import "testing"

func TestSoma(t *testing.T) {
	teste := Soma(10, 10)
	resultado := 20

	if teste != resultado {
		t.Error("expected:", resultado, "got:", teste)
	}
}
// Esse da erro!
func TestSoma2(t *testing.T) {
	teste := Soma(10, 11)
	resultado := 20

	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMulti(t *testing.T) {
	teste := Multiplicacao(10,10)
	resultado := 100

	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

// Esse da erro!
func TestMulti2(t *testing.T) {
	teste := Multiplicacao(10,11)
	resultado := 100

	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}