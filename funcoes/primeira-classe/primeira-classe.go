package main

import "fmt"

/* Conceito de Primeira Classe:

É um conceito fundamental, muito importante dentro de GO, isso permite um leque de possibilidades, principalmente pra programação de paradigma funcional, antigamente não era tão evidente, por que era um padrão que demandava uma quantidade de memória muito maior, pois não trabalha com referência de memória, trabalha com cópias de valores, e evolução desse valor. Essa questão de evoluir um dado e não alterar um dado já existente, demanda mais memória, mas hoje temos capacidade de memória muito maior, do que quando esse paradigma surgiu, o paradigma funcional, surgiu antes do paradigma de Orientação ao Objeto.

*/

// Defino uma variável que armazena uma função anônima.
var soma = func(a, b int) int {
	return a + b
}

// Posso armazenar uma função numa variável qualquer e invocar essa variável em qlqer parte do meu sistema, passando o argumento dentro dela.
func main() {
	fmt.Println(soma(2, 3))
	// := anotação reduzida
	subtracao := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtracao(4, 2))
}
