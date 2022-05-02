package main

import "fmt"

// exercicio:

func notaParaConceito(nota float64) string {
	switch {
	case nota > 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 2 && nota < 5:
		return "D"
	case nota >= 0 && nota < 2:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(-9))
	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(0))
}
