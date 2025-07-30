package main

import "testing"

type test struct {
	data []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},

		test{[]int{2, 2}, 4},

		test{
			[]int{1, 1},
			2,
		},
	}

	for _, v := range tests {
		valorDaSoma := Soma(v.data...)
		if valorDaSoma != v.answer {
			t.Error("expected:", v.answer, "got:", valorDaSoma)
		}
	}
}

// func TestSoma(t *testing.T) {
// 	teste := Soma(10, 10)
// 	resultado := 20

// 	if teste != resultado {
// 		t.Error("expected:", resultado, "got:", teste)
// 	}
// }

// func TestMulti(t *testing.T) {
// 	teste := Multiplicacao(10,10)
// 	resultado := 100

// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }