package main

import "fmt"

// struct = estrutura

/*
	Dentro da estrutura, eu vou colocar todo tipo de dado que tem nessa estrutura, na verdade é um agrupador de dados...

	Parece uma classe no POO, mas não tem herança
*/

// identificador - tipo
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// conceito receiver -> quero aplicar essa função a qual receptor, quem será o dono dessa função, para que eu possa posteriormente, para que eu possa fazer algo do tipo produto.algumacoisa, um método por cima de alguma estrutura.

// criar um método ou uma função que esta associada a esta estrutura.

// Método: função com receiver(receptor)

// dentro do argumento, coloco quem vai ser o receptor, o dono da função, nesse caso vai ser o produto, que dentro ali do argumento referencio como a letra p, depois com a nomenclatura normal, o identificador(nome) da funcao, os parametros se houver, e o tipo do retorno se a função tiver.
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

var produto1 produto

func main() {
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	// Outra forma de criar uma struct, de forma mais reduzida
	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
