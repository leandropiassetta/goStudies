package main

import (
	"fmt"
	"time"
)

// exercicio:

func notaParaConceito(n float64) string {
	nota := int(n)
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
	t := time.Now()     // hora atual
	switch /* true */ { // switch true --> Se não colocar o switch vinculado a algum caso, ele vai procurar algum case, que seja verdadeiro, ai irei botar alguns cases com verdadeiros ou falso e o primeiro verdadeiro que ele encontrar, ele vai executar..
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}

	fmt.Println(notaParaConceito(-9))
	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(0))
}
