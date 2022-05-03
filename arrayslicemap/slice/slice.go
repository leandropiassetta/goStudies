package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array - com numero dentro do colchete array com tipo fixo
	s1 := []int{1, 2, 3}  // slice - não tem número dentro do colchete e é do tipo variável
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1)) // pacote pra ver os tipos das variáveis

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3] /*vai até o indice 2*/ // não gera um array diferente, é uma estrutura que aponta pro primeiro elemento que ele aponta, é uma estrutura continua.
	fmt.Println(a2, s2)

	// vai do inicio da estrutura do array e vai até o indice 1.
	s3 := a2[:2] // novo slice, mas aponta para o mesmo array.
	fmt.Println(a2, s3)

	// voce pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array.

	s4 := s2[:1] // posso fazer slice de um slice...
	fmt.Println(s2, s4)
}
