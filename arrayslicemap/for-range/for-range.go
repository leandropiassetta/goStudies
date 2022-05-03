package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta! [...]É uma anotação que diz que é um array, ele tem um número fixo, mas quero que você conte a partir da quantidade de elementos que eu inicializar nesse array.

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// range parece o forEach do JS

	// _ --> dessa forma ignora o indíce.
	for _, num := range numeros {
		fmt.Println(num)
	}
}
