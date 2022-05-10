package main

import (
	"fmt"
	"time"
)

// Usando goroutine e channel

// channel é um tipo...

// Channel (canal) - é a forma de comunicação entre goroutines, é um ponto de sincronismo pra você receber esse dado e a partir dali, de todos os dados que vieram de vários processamentos independentes, você tem como continuar o seu programa.

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	// espera um segundo
	// depois envia pro canal o resultado de 2 * base
	// e assim sucessivamente com as outras linhas de códigos abaixo
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

/*
A função interna make aloca e inicializa um objeto do tipo slice, map ou chan (somente). Como novo, o primeiro argumento é um tipo, não um valor. Ao contrário de new, o tipo de retorno de make é o mesmo que o tipo de seu argumento, não um ponteiro para ele. A especificação do resultado depende do tipo:
*/
func main() {
	c := make(chan int)
	// A partir daqui essa função vai executar de forma independente

	// Depois que é startado ele continua a execução da função main
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // Os primeiros dois valores que eu receber, vai ser armazenado na variável a e b
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println(<-c) // O último dos três valores eu recebo aqui.
}
