package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// Não defino o tamanho do array é um slice
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	// Passando um slice para uma função variática, ele espalha, esse elementos para a função
	imprimirAprovados(aprovados...)
}
