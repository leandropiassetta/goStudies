package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Criar um pequeno bloco de inicialização de uma variável antes de executar um if e essa variável vai ficar disponível, o escopo dela vai ser apenas o bloco if/else.

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // Pacote rand(random), passar pra ele um valor aleatorio, e pegar um nano segundo e passar como fonte pra ele gerar um numero aleatorio;
	r := rand.New(s)                           // sempre que executar o  código ele vai gerar um numero novo

	// Ele vai pegar o nano segundo atual quando eu executar o código
	return r.Intn(10) // ele vai gerar um número aleatório até 10
}

func main() {
	// Eu vou definir um bloco if, só que associado ao bloco if, vou colocar uma sentença de código pra inicializar uma variável que estará presente apenas dentro do bloco if e ao bloco else associado ao bloco if

	// obs: parecido com o for
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}
