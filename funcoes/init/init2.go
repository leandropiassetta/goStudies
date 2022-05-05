package main

import "fmt"

// Função init:
// Eu posso ter um init para cada arquivo, por convenção ele vai executar este método
/*
	Existe uma convenção tb usada no GO, uma funçao init, essa função por padrão é executada, quando aquele arquivo em Go é interprertado.
*/

func init() {
	fmt.Println("Inicializando....2")
}

// Eu nao posso ter dois métodos main, dentro de um mesmo código...

// P/ executar esse arquivo entro até a pasta dele, e rodo o comando go run *.go
