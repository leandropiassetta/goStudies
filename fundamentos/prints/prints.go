package main

import "fmt" // pacote fmt entrada e saida, formatada.

func main() {
	fmt.Print("Mesma") // Coloca na mesma linha aquilo que voce colocou dnetro da funcao Print, mesmo que seja duas sentença de codigo diferente
	fmt.Printf(" linha.")

	fmt.Println(" Nova") // O Println imprime o texto e no final cria uma nova linha.
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x), dessa forma nao funciona, nao permite concatenação de um valor string com um valor numérico diretamente
	xstring := fmt.Sprint(x) // essa função Sprint, retorna uma string
	fmt.Println("O valor de x é " + xstring)
	fmt.Println("O valor de x é", x) // dessa forma consigo fazer a concatenação

	fmt.Printf("O valor de x é %.2f", x) // tem mais poder se usar o Printf, diminuir as casas decimais, e dentro da STRING, posso colocar ganchos, estilo template literals no JS, '%.2f' pega apenas duas casas decimais.

	a := 1
	b := 1.999
	c := false
	d := "opa"
	/*%d -> INTEIRO %.1f -> FLOAT com uma casa decimal. %t -> BOOL %s -> STRING*/
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // fica na mesma linha, se quiser uma nova linha, usa o \n -> new line
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // %v --> imprime de uma forma correta, pra varios tipos diferente de uma forma mais generica
}
