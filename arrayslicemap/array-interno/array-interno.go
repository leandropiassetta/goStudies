package main

import "fmt"

// Slice -> O mesmo array interno é capaz de ser referenciado para vários slices, esse é o ganho do slice, é uma estrutura extremamente leve, que eu posso com a mesma base, com o array estatico embaixo, eu posso pegar trechos diferentes e ter uma estrutura flexivel para utilizar.
func main() {
	// Os dois slices estão apontando para o mesmo array interno
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)
}
