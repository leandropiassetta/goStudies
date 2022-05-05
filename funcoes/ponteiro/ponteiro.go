package main

import "fmt"

// Em Go por padrão você passa cópias de valores, o que é bom por que gera menos efeitos colaterais na aplicação, mas você pode passar tb uma referência pra um objeto para que um método possa alterar essa referência, mexer diretamente no dado que você como pasosu como parâmetro....

func inc1(n int) {
	n++ // n = n + 1
}

// revisão: Um ponteiro é representado por um *

func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta

	*n++ // o * retorna um valor e eu consigo incrementar aquele valor, que aquele ponteiro faz referência
}

func main() {
	n := 1

	inc1(n) // por valor, a função opera com a cópia do valor, o que faz com que a variavel n, nao seja afetada...
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // por referência

	fmt.Println(n)
}

// Priorize passagens por valor e evite passar por referências.
