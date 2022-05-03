package main // programa executável

import "fmt"

func main() {
	// Array = homogênea (mesmo tipo) e estática (tamanho fixo), não varia com o tempo, durante todo o ciclo de vida desse array, a não ser se voce pegar uma variavel e apontar para uma outra estrutura de tamanho diferente, mas aquele array vai sempre permanecer com aquele tamanho.

	// Array é uma estrutura indexada, [pos0,pos1, pos2]
	// pegar o tamanho do array = len(array)

	var notas [3]float64 //[3] tamanho do array, estrutura estática ou seja fixa
	// Por padrão ele usa os 0 do float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10 -> erro de compilação, o indíce 3 é um indice inválido, nao existe.
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media) // Printf -formatar a mensagem com o %, .2f pra ser do tipo float com duas casas decimais.
}
