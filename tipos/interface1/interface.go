package main

import "fmt"

//*POO / Polimorfismo => o principal conceito do polimorfismo é a propriedade de duas ou mais classes derivadas (Macaco, Homem e Baleia) de uma mesma superclasse (Mamifero) responderem a mesma mensagem (locomoverSe()), cada uma de uma forma diferente (pulando de galho em galho, andando e nadando, respectivamente).

// INTERFACE

/*
	Em GO, se você tem uma estrutura que respeita os métodos de uma interface, mesmo que você não diga nada, o próprio compilador infere dizendo que esta estrutura tem métodos que suportam os métodos da interface, automaticamente ela é daquele tipo.
	Ai posso usar o polimorfismo, sem ter que explicitamente ter que amarrar na struct dizendo que ela respeita interface tal, que ela implementa interface tal, porque automaticamente o compilador checa e ele sabe que determinada estrutura pode ser passado de forma polimorfica para uma interface ou não.
*/

type imprimivel interface {
	// Eu vou definir um método que toda interface imprimivel tem que ter, definindo o método toString(), que retorna uma string, para você saber se uma estrutura é imprimível ou não, ela precisa ter esse método associado a ela.
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// Interfaces são implementadas implicitamente

// receiver pessoa e estou botando um alias p
func (p pessoa) toString() string {
	// concatenando
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	// Sprintf formata de acordo com um especificador de formato e retorna a string resultante.
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// recebe como parâmetro qualquer coisa que respeita a interface imprimivel..
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// qualquer coisa que tenha o tipo imprimivel e irei atribuir a uma pessoa.

	// Polimorfismo = Múltiplas formas
	// uma hora a variável coisa tem forma de pessoa, outra tem forma de produto..
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// imprimir(produto{"Calça Jeans", 79.90})

	p2 := produto{"Banana", 7}
	imprimir(p2)
}
