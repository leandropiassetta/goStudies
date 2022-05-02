package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A" // mesmo sem return ele executaria uma dessas linhas
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
	// Se não entrar em nenhuma daquelas linhas acima da estrutura if/elseif/else, ele continuaria executando e entraria aqui...
}

// SWITCH -> desenhado para ter multiplas seleções
func main() {
	fmt.Println(notaParaConceito(9.9))
	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(3.9))
	fmt.Println(notaParaConceito(2.9))
}
