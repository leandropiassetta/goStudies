package main

import "fmt"

// Map -> conceito de estrutura chave e valor, estilo objeto em JS e dicionário(dict) em Python...

func main() {
	// chave do map é do tipo int e o valor é string
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678901] = "Maria"
	aprovados[98927485345] = "Pedro"
	aprovados[72615384040] = "Ana"

	fmt.Println(aprovados)

	// for em map -> primeiro argumento percorre  a chave, o segundo o valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	// Para ler um valor de um map

	fmt.Println(aprovados[72615384040])
	// Excluir um valor de um map
	// o primeiro parâmetro é o map que voce quer utilizar e o segundo é o valor.
	delete(aprovados, 72615384040)
	fmt.Println(aprovados[72615384040])
}
