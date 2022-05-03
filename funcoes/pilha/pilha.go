package main

import "runtime/debug"

func f3() {
	debug.PrintStack() // pacote nativo do Go, debug, usando o método que PrintStack que vai imprimir a pilha de execução de um programa no momento que essa função for chamada.
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
