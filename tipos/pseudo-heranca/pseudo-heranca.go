package main

import "fmt"

// Imitando o conceito de herança de POO

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anônimos, tudo que é de carro, vai ficar disponiverl diretamente dentro de ferrari, vai da a sensação que estou herdando da estrutura de carro, mas na verdade é uma composição...
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
