package main

import "fmt"

// Tipo Personalizado:

/*
	Temos a possibilidade de criar tipos personalizados, a partir de um tipo base eu posso criar um tipo, temos a possibilidade de associar novos métodos nesse tipo que eu personalizei como receiver, como método que eu quero colocar.
*/

type nota float64

// funçao que recebe como receiver uma nota
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(5.8))
	fmt.Println(notaParaConceito(4.8))
	fmt.Println(notaParaConceito(1.8))
}
