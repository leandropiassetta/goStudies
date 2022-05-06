package main

import "fmt"

// Composição de interfaces:

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// composisção das duas interface
type esportivoLuxuoso interface {
	esportivo
	luxuoso

	// pode adicionar novos métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo....")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	// Ele executa os dois métodos a partir de uma interface que é um composição de duas interface...
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
