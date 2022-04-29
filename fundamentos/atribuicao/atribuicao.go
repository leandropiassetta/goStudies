package main

import "fmt"

func main() {
	var b byte = 6
	fmt.Println(b)

	i := 3 // inferência de tipo

	i += 3 // i = i + 3 ->  atribuição somativa
	i -= 3 // i = i - 3 -> atribuição subtrativa
	i /= 2 // i = i / 2 -> atribuição divisiva
	i *= 2 // i = i * 2 -> atribuição multiplicativa
	i %= 2 // i = i % 2 -> atribuição por módulo

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x // trocar os valores de duas variáveis
	fmt.Println(x, y)
}
