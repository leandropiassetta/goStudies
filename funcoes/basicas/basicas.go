package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

// parâmetro e o tipo do parâmetro
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// a função retorna uma string
func f3() string {
	return "F3"
}

// parametros do mesmo tipo, posso declarar o tipo dela apenas uma vez
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) // a função Sprintf, nao imprime no console, e sim retorna uma string formatada.
}

// Go: te dá a opção que uma função te retorna múltiplos valores.

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5", r51, r52)
}
