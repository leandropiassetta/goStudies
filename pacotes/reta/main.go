package main

import "fmt"

// Eu consigo usar outras funções por causa dos elementos que são visiveis pra fora do pacote, quando coloquei letra maiusculas nela.

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	// Eu consigo usar a função catetos, pq ele é privado no nivel de pacote e nao no nivel do arquivo

	// Eu preciso executar dentro do console.
	// cd pacote/reta
	// go run *.go
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
