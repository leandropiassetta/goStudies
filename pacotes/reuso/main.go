package main

import (
	"fmt"

	"github.com/leandropiassetta/area-func"
)

func main() {
	fmt.Println(area.Circ(4.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// Essa função não existe, nao foi definida, ela foi definida de forma privada, dentro de um pacote area, se eu estou em outro pacote não consigo utilizar essa função.
	// fmt.Println(area._TrianguloEq(5.0, 3.0))
}
