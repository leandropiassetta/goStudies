package main

import (
	"fmt"
	"time"
)

// recebe um channel de inteiros
func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	// ele só vai prosseguir no canal, quando alguem do outro lado, ler esse dado, vai mandar e alguem vai ler, e ai vai pro proximo dado do canal....
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock:  Ele sabe que não tem ninguem que possa enviar esse dado pro canal, pq todas as goroutines ja morreram a essa altura do campeonato.
	fmt.Println("Fim")
}
