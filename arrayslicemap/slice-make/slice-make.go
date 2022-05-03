package main

import "fmt"

func main() {
	// numero 10, significa que vai ter 10 elementos.
	s := make([]int, 10) // cria um slice a partir do método make
	s[9] = 12
	fmt.Println(s)

	// 20 -> o array interno desse slice vai ter 20 elementos
	s = make([]int, 10, 20)
	// cap -> pegar a capacidade máxima desse slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
