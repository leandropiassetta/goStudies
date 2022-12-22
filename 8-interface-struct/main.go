package main

import "fmt"

type fruta interface {
	// contrato
	nome() string
}

type melancia struct{}

func (m melancia) nome() string {
	return "mlenacia"
}

func vitaminaD(f fruta) {
	fmt.Println("Vitamina de", f.nome())
}

type banana struct{}

func (b banana) nome() string {
	return "banana"
}

func main() {
	var m fruta
	m = melancia{}
	nome := m.nome()
	fmt.Println(nome)

	vitaminaD(m)

	var b fruta
	b = banana{}
	nome = b.nome()
	fmt.Println(nome)

	vitaminaD(b)
}
