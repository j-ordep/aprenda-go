package main

import "fmt"

type veiculo struct {
	nome   string
	portas string
	cor    string
}

type caminhonete struct {
	tracaoNasQuarto bool
	tipo            veiculo
}

type sedan struct {
	modeloLuxo bool
	veiculo
}

func main() {

	carro1 := veiculo{"corsa", "2 portas", "vermelho"}

	fmt.Println(carro1.nome, carro1.cor, carro1.portas)
	
	carro2 := caminhonete{true, veiculo{"hilux", "4 portas", "azul"}}
	
	fmt.Println(carro2.tipo.nome, carro2.tipo.cor, "tração nas 4 rodas?",carro2.tracaoNasQuarto)
	
	carro3 := sedan{false, veiculo{"hb20", "4 portas", "branco"}}
	
	fmt.Println(carro3.nome, carro3.cor, "modelo de luxo?",carro3.modeloLuxo)

}