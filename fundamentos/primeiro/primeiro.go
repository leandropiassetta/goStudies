// Programas executáveis iniciam pelo pacote main

package main // pacote principal do programa é o main por que esse programa que irá ser executado

/*
Os códigos em Go são organizados em pacotes e para usá-los é necessário declarar um ou vários imports
*/

import "fmt" // import de outro um outro pacote de um core da linguagem

// função main que é a porta de entrada de um programa GO
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
}
