package main

import "fmt"

// uint -> só permite números inteiros sem sinal, um inteiro positivo
func fatorial(n uint) uint {
	switch {

	// condição de parada
	case n == 0:
		return 1
	default:
		fmt.Println(n)

		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
