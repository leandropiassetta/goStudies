package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

// Função exec, vai receber como primeiro parâmetro uma funcao, alem de dizer o tipo func, tenho que tb dizer o tipo de valor que ela vai receber como parâmetros, além do tipo de valor que a função irá retornar.
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 2, 2)
	fmt.Println(resultado)
}
