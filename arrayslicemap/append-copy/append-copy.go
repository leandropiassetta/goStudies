package main

import "fmt"

func main() {
	// Array em Go tem tamanho fixo
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	/*o primeiro argumento do append precisa ser um slice*/
	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// vou pegar tudo que tem dentro do slice1 e copiar pro slice2
	copy(slice2, slice1)
	// no copy, nao pode passar array, tem que ser string ou slice
	/* 1º argumento recebe a copia do 2º argumento do método do copy */
	fmt.Println(slice2)
}
