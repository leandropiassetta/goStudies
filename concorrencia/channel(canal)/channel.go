package main

import "fmt"

// channel é um tipo, dentro de um channel, pode trafegar dados de um determinado tipo...
func main() {
	// criando um canal que dentro desse canal de comunicação, vai estar sendo enviados valores inteiros.
	// segundo parâmetro é um buffer
	ch := make(chan int, 1)

	// estou enviando para um canal o valor 1
	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	// channel é um mecanismo de sincronismo
	fmt.Println(<-ch)
}
