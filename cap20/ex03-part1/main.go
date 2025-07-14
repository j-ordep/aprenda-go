package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Utilizando goroutines, crie um programa incrementador:
// Tenha uma variável com o valor da contagem
// Crie um monte de goroutines, onde cada uma deve:
// Ler o valor do contador
// Salvar este valor
// Fazer yield da thread com runtime.Gosched()
// Incrementar o valor salvo
// Copiar o novo valor para a variável inicial
// Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// Demonstre que há uma condição de corrida utilizando a flag -race

var wb sync.WaitGroup

func main() {
	contador := 0

	maxGoroutines := 100

	wb.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func ()  {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println(contador)
			wb.Done()
		}()
	}

	wb.Wait()

	fmt.Println("total:", contador)
}