package main

import "fmt"

// GO é uma linguagem fortemente tipada, uma vez que voce define um tipo pra uma variavel, voce tem que seguir aquele tipo até o final do seu programa.. Mas isso nao te impede de ter um tipo mais genérico que pode flexibilizar o uso de variaveis e podemos fazer isso a partir do uso da interface como um tipo..

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)
	// Você pode aplicar multiplos tipos dentro desse tipo mais generico dentro desse tipo mais genérico que seria uma interface vazia...
	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)
}
