package main

import "fmt"

// Funcoes variaticas -- que recebem quantidade de parâmetros variáveis
// JS ->> operador spread e rest

func media(numeros ...float64) float64 {
	// Dentro da função irei trabalhar esses parâmetros como array
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	// como o total é um float64 e o tamanho é um inteiro, converto o tamanho em float64
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7, 9, 8.8, 10))
}
