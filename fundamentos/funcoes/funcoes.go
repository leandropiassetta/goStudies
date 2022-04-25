package main

import "fmt"

// nome da variavel e tipo, nome da variavel e tipo, e retorna o tipo do resultado que a função irá retornar.
func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) /*como nao retorna nada nao preciso explicitar o tipo do valor do retorno*/ {
	fmt.Println(valor)
}
