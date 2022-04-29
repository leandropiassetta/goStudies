package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 { // não tem parenteses em Go no escopo do if
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	} // tem que ter definição do bloco..
}

// função main porta de entrada no programa Go
func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
