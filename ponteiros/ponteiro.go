package main

import "fmt"

// func main() {
// 	fmt.Println("Ponteiros")

// 	var variavel1 int = 10

// 	var variavel2 int = variavel1 // copia do valor

// 	fmt.Println(variavel1, variavel2)

// 	// PONTEIRO É UMA REFERENCIA DE MEMORIA

// 	var variavel3 int
// 	var ponteiro *int

// 	fmt.Println(variavel3, ponteiro)

// 	variavel3 = 100
// 	ponteiro = &variavel3

// 	fmt.Println(variavel3, *ponteiro) // desreferenciação, vejo o valor que esta armazenado no ponteiro

// }

func main() {
	i, j := 42, 2701

	p := &i         // endereço de memória para i (ponteiro de i)
	fmt.Println(*p) // lê o valor de i através do ponteiro
	*p = 21         // altera o valor de i através do ponteiro
	fmt.Println(i)  // lê o novo valor de i

	p = &j         // ponteiro de J
	*p = *p / 37   // divide o valor de J usando o ponteiro
	fmt.Println(j) // verifica o novo valor de J
}
