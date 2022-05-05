package main

import "fmt"

// Conceito de Recursividade:

/*
	É uma função que chama a si própria, uma das questões mais importantes para uma função recursiva, é que ela tenha uma condição de parada, para que ela nao fique se chamando de uma forma indefinida, gerando uma pilha gigante, uma stack de execução muito grande, gerando o famoso stackoverflow, o estouro de pilha.
*/

// É muito comum em funções GO, funções que retornam um valor e também retornam um erro.
func fatorial(n int) (int, error) {
	switch {
	// Quando o n é menor que zero, eu nao posso ter um fatorial de numero negativo, isso é um erro, obrigatoriamente tenho que retornar o valor, o -1 é pra dizer que é um valor inválido
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	// Por definição o fatorial de 1 e 0, os dois representam o resultado 1
	case n == 0 || n == 1:
		return 1, nil // o resultado é um, e o erro é nil(null), não vou ter erro associado.
	// no caso default eu vou chamar de forma recursiva.
	// vai chamando a funçao, diminuindo ela de 1 em 1
	// 3! 3*2*1 = 6
	// 4! 4*3*2*1 = 24
	default:

		fatorialAnterior, _ := fatorial(n - 1)

		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	// ignorando o erro, pq sei que nao vai dar erro.
	fmt.Println(resultado)
	// ignorando o resultado, pq sei que nao vai ter resultado.
	_, err := fatorial(-4)
	// forma muito comum de você tratar o erro.
	if err != nil {
		fmt.Println(err)
	}
}
