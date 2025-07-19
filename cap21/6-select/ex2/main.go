package main

import "fmt"

// Func 1 recebe X valores de canal(receive), depois manda qualquer coisa pra chan quit(send)
// Func 2 for infinito, select: case envia pra canal(send), case recebe(receive) de quit

/*
Este programa usa 2 canais: 'canal' e 'quit'.

- 'canal' é usado para enviar números continuamente.
- 'quit' é um canal de controle para sinalizar quando o envio deve parar.

Fluxo geral:

1. A função enviaProCanal() envia números continuamente no canal, usando um select:
   - Enquanto houver alguém consumindo do canal (recebeQuit), ela continua enviando valores.
   - Se receber um sinal no canal quit, ela encerra.

2. A função recebeQuit() consome 50 valores do canal.
   - Após consumir 50 valores, ela envia um valor pelo canal quit, avisando a função enviaProCanal() para parar.

O select:
- Fica escutando os dois canais ao mesmo tempo.
- Só executa um dos cases de cada vez:
   - Se conseguir enviar no canal (case canal <- qualquerCoisa), envia e continua.
   - Se receber algo no canal quit (case <-quit), encerra o loop.

Conclusão:
- Enquanto recebeQuit() estiver lendo, enviaProCanal() continua enviando.
- Quando recebeQuit() terminar, enviaProCanal() para automaticamente.
*/

func main() {

	canal := make(chan int)
	quit := make(chan int) // sinal

	go recebeQuit(canal, quit)
	enviaProCanal(canal, quit)

}

func recebeQuit(canal, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido: ", <-canal) // recebe o valor do canal até o loop chegar ao fim, quando sai do loop quit recebe 0 (quit <- 0)
	}

	quit <- 0 // sinal
}

func enviaProCanal(canal, quit chan int) {
	qualquerCoisa := 1
	for {
		select {
			case canal <- qualquerCoisa: // fica enviando esse valor ao canal enquanto alguem estiver recebendo esse valor, no caso a func recebeQuit()
				qualquerCoisa++
			case <-quit: // aqui ele retira o valor de <-quit, porem ele so retira se alguem enviar, e só sera enviado quando o for-i da func recebeQuit() terminar
				return
		}
	}
}