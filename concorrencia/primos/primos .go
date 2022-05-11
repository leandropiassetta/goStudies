package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// n = quantidade de primos que eu quero retornar
// c = valores primos serao enviados para o canal
func primos(n int, c chan int) {
	inicio := 2 // controle
	for i := 0; i < n; i++ {
		/*sem condição de parada*/
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo // enviando dados para o canal (escrita)
				inicio = primo + 1
				time.Sleep(time.Millisecond * 280)
				break
			}
		}
	}
	close(c)
	// se não colocar a funcao close, gera um deadlock
	// finaliza o laço for, dentro desse canal de comunicação não vai ser enviado mais nenhum dado..
}

func main() {
	// Dentro desse canal de inteiro, vai ter um buffer com 30 posições
	c := make(chan int, 30)
	// cap = capacidade que tem o canal = 30
	// sempre que chegar novos valores esse laço vai interando
	go primos(cap(c), c)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
