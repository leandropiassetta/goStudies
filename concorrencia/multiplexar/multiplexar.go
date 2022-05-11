package main

import (
	"fmt"

	"github.com/leandropiassetta/go-html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		// sempre que chegar dados no canal de origem, eu vou encaminhar ele pro canal de destino.
		// enquanto não tiver dados no canal de origem, esse for fica bloqueado..
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	// recebe o dado num parametro e encaminha pro canal c, recém criado..
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c)
}
