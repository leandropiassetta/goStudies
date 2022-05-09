package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second) // significa que ele vai passar um segundo parado esse código
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Essa forma de chamar função é uma forma serial..
	// fale("Maria", "Pq você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// Ao botar essa palavra go antes de um método, significa que ele vai criar uma goroutine, ele vai executar o código de uma forma indepedente, seria como se fosse uma thread
	// thread -> abrir uma linha de execução independente.. executando em paralelo ou em forma concorrente.

	// go fale("Maria", "Ei....", 500)
	// go fale("João", "Opa....", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim!")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
}
