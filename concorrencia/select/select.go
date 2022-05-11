package main

import (
	"fmt"
	"time"

	"github.com/leandropiassetta/go-html"
)

// SELECT: estrutura de controle especifica pra trabalhar com concorrencia dentro de GO

// estruturas de controle: for, if/else, switch, select;

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Estrutura de controle específica para concorrência
	// Primeiro dado no canal que chegar, ele vai executar aquele case e ai retorna a execução, executa todas as sentenças que tiverem associada aquela senteça e sai do laço para continuar a executar os codigos restantes..
	select {
	// caso t1 receba o valor de channel1, retorna t1
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
		// timeout, depois de um determinado valor, se nao chegar, retorna "todos perderam" pq demorou demais
		// Se estourar o tempo ele retorna o resultado do timeout pro canal.
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!!!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://wwww.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
