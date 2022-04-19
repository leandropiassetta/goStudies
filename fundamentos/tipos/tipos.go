package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//----------------------------1byte(8bits) 2byte(short) 4byte(int) 8byte(long)
	// sem sinal (só positivos)... uint8       uint16       uint32    uint64
	// u(unsigned)int = como o range dele é só positivo ele vai ocupar todos os bytes

	var b byte = 255 // byte é um apelido pra uint8

	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	// math.MaxInt64 -> vai pegar o maior valor possivel de numero inteiro com 64 bits
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representação um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // por padrao o literal é float64

	// boolean

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string

	s1 := "Olá meu nome é Leandro"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) // funções que trabalham em cima de dados.

	// string com multiplas linhas

	s2 := `Òlá 
	meu
	nome
	é
	Leandro`
	fmt.Println("O tamanho da string é", len(s2))

	// char ???
	char := '9'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
