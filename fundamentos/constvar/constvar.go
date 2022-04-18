package main // define o pacote o main

import (
	"fmt"
	m "math"
)

func main() { // e a função main para que esse codigo seja executavel

	// Para definir a variável ou constante, escrevo o nome da constante e o tipo dessa constante;
	const PI float64 = 3.1415
	raio := 3.2
	// tipo(float64) inferido pelo compilador;

	// area = escrita dessa forma significa que a variavel area ja foi definida e eu estou  atribuindo um valor pra ela.

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2) // dessa forma eu estou definindo essa variavel, atribuindo um valor e inicializando essa variavel.

	// outra forma de definir uma constante
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
	// numa unica sentença de codigo, voce pode atribuir valor a varias variaveis

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)

	fmt.Println("A área da circunferência é", area)

	// Em Go, eu nao sou obrigado a colocar o tipo da variavel se eu atribuo um valor pra ela, pq o compilador consegue fazer a inferencia a ela..
}

// OBS: Em Go se você define uma variável voce obrigatoriamente tem que usar essa variável se nao o propgrama nao irá funcionar, não vai compilar o código.
