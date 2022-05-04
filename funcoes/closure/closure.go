package main

import "fmt"

/*
	Função Closure:

	Closure -> Algo que envolve alguma coisa.

	A função que envolve outra função, isso faz com que a função interna que seja envolvida pela externa, ela tem memoria e ciencia, ela entende o local aonde foi declarada, ela entende em que contexto léxico o contexto da palavra que ela foi declarada, o contexto léxico é o local fisico no qual você declarou alguma coisa.
*/

func closure() func() {
	x := 10
	funcao := func() {
		fmt.Println(x)
	}
	return funcao
}

// Go suporta função closure, ele vai ter essa memória das suas origem.

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
