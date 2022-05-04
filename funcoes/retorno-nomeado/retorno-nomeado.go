package main

import "fmt"

// os dois parametros sao do mesmo tipo, posso botar o valor do tipo só no final do argumento da função.
// posso definir nomes para os retornos.
func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
