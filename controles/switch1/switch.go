package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	// Switch não tem uma expressão que retorna verdadeiro ou falso
	switch nota { // No GO, se ele entra no case 10, ele sai do bloco switch.
	case 10:
		fallthrough // palavra reservada, vai continuar executando os proximos blocos mesmo encontrando o case.
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11.1))
}
