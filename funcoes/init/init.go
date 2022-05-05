package main

import "fmt"

// Função init:

/*
	Existe uma convenção tb usada no GO, uma funçao init, essa função por padrão é executada, quando aquele arquivo em Go é interprertado.
*/

func init() {
	fmt.Println("Inicializando....")
}

func main() {
	fmt.Println("Main....")
}
