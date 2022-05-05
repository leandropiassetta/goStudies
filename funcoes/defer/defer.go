package main

import "fmt"

/*
	Função Defer:

	Em Go, temos uma função reservada da linguagem chamada defer, que é uma função que ela atrasa, digamos assim uma execução, botamos essa palavra reservada antes de uma sentença de código, quando você quer atrasar essa execução dessa sentença de código, até o momento antes de retornar do método sair e chamar o return.

	Por exemplo,  você abre uma conexão com o banco de dados e você quer lembrar de fechar esse recurso, antes de sair do método, abre uma conexão, na linha seguinte, coloca a palavra defer e a sentença de codigo que fecha a conexao, e voce garante que aquela sentença de codigo vai ser atrasada ate momento antes do retorno da funcao, atrasa uma determinada execução.

	Obs: lembra o async/ await do JS

*/

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!") // embora seja a primeira linha ele vai ser executado logo antes de retornar o valor do método, mesmo que eu tenha mais de um tipo de fluxo dentro do meu método.

	// Eu disse, atrasa a execução disso pro ultimo momento posssivel dentro desse método.
	if numero > 5000 {
		fmt.Println("Valor alto....")
		return 5000
	}
	// como eu tenho um return dentro do if, nao preciso do else..
	fmt.Println("Valo baixo....")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3300))
}
