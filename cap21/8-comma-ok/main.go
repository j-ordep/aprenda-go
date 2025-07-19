package main

import (
	"fmt"
)

/* Problema!

	começa com 0 par e termina com 0 inpar

	Utilizando comma ok para resolver esse problema


	IMPORTANTE: Enviar o sinal pelo canal 'quit' ANTES de fechar os canais 'par' e 'inpar'.

	Motivo:
	- A função receiveNum() está em um loop infinito, esperando valores de 'par', 'inpar' ou 'quit'.
	- Se os canais 'par' e 'inpar' forem fechados antes do envio no 'quit', o loop continuará tentando ler deles.
	- Quando um canal é fechado e tentamos ler dele, o Go devolve o zero value (no caso, 0), sem bloquear.
	- Por isso, continuaria imprimindo "0 é par" ou "0 é inpar" indefinidamente.

	Solução:
	- Enviar o sinal pelo canal 'quit' ANTES de fechar 'par' e 'inpar'.
	- O select detecta o recebimento no canal 'quit' e executa o retorno (return), encerrando o loop imediatamente.
	- Após o envio do sinal, podemos fechar os canais com segurança, pois o loop já parou.

	Resumo:
	- quit <- true serve como sinal de parada.
	- Enviar esse sinal ANTES de fechar os canais previne leituras de canais fechados (que causariam valores zero).

*/

func main() {
	par := make(chan int)
	inpar := make(chan int)
	quit := make(chan bool)

	go sendNum(par, inpar, quit)
	receiveNum(par, inpar, quit)

}

func sendNum(par, inpar chan int, quit chan bool) {
	for i := 0; i <= 30; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			inpar <- i
		}
	}
	quit <- true // quit deve vir antes para enviar o sinal e encerrar o loop, se não ele continua rodando o loop mesmo com o canal fechado
	close(par)
	close(inpar)
}

func receiveNum(par, inpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println(v, "é par")

		case v := <-inpar:
			fmt.Println(v, "é inpar")

		case v, ok := <-quit:
			if !ok {
				fmt.Println("Comma não ta ok", v)
			} else {
				fmt.Println("Encerrando, recebemos comma ok =", v) // nesse caso quit é sempre true, então comma ok sempre seré true
			}
			return
		}
	}
} 

