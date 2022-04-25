package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // você precisa explicitamente converter um tipo inteiro para um tipo float, para nao ter nenhum problema quando calcular

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // pra concatenar convertendo string = strconv, Itoa = int to string

	// string para int

	num, _ := strconv.Atoi("123") // retorna duas coisas, o primeiro valor é um numero o segundo é um erro, caso voce passe uma string  que nao é um inteiro, Atoi= converte String em valor numérico.
	fmt.Println(num - 122)

	// conversão tipo boolean
	// 0 = false
	// todos os numeros diferente de 1, é falso!
	// 1 = true
	b, _ := strconv.ParseBool("true")
	// estrutura de controle
	if b {
		fmt.Println("Verdadeiro")
	}
}
