package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) /*retorna 3 bool*/ {
	comprarTv50 := trab1 && trab2 // operação binária pq usa dois operandos.
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete /* operador unnário*/)
}
