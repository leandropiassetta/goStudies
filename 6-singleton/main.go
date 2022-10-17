package main

/* O que é singleton?
"Garante a criação de uma única instância de uma classe através de toda a aplicação, mantendo um ponto global de acesso ao seu projeto."

SINGLETON: Motivação através de exemplos..

* Única conexão com banco de dados.
* Acesso ao sistema de arquivos de um sistema.
* Sistema de logs

* Lazy construction

* Não há parâmetros para sua construção
* Acessível de forma global

*/

/* SINGLETON: ESTRUTURA PRINCIPAL

* Método construtor é privado
* Método "getInstance" responsável por criar o objeto
	- Ele verifica se o objeto existe, caso contrário cria
* Possui outros métodos como uma classe qualquer

*/

/*
	NOT THREAD SAFE VS THREAD SAFE

	* PADRÃO "ORIGIAL" singleton não é thread safe. Isso é grave
	*	Existem formas de contornar o problema atraves de variações desse padrão.

*/

/*
	SINGLETON COMO UM ANTI-PATTERN

	* ALTO ACOPLAMENTO
	* Acesso global de um objeto
	* Não é thread safe
	* Dificuldade para trabalhar c/ testes
	* QUEBRA O SOLID
		- SRP(ela possiu a responsabilidade de se gerenciar)
	* Pode ser facilmente substituida através de injeção de dependência
	(DI)

*/

func main() {

}
