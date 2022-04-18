package main

import "fmt" // Aplica a formatação tradicional no código

func main() {
	// Printf -> em teoria ele vai substituir a porcentagem, por algum valor personalizado, como se fosse o template strings ou literals do JS
	fmt.Printf("Outro programa em %s!", "Go") // Ai eu altero o que esta na porcentagem. colocando um segundo argumento dentro do método Printf, que pertence ao pacote fmt.
}

/*
	COMANDOS NO TERMINAL
go -> O comando go, mostra varias possibilidades do uso do GO

go help -> Pode chamar o go help, passando outro comando, e ai ele explicará o que esse comando pode fazer, por ex: go help get, o comando get faz o download de pacotes, a partir do caminho de importação, posso baixar pacotes direto do github.

go version -> mostra a versão do GO

godoc -http=:6060 -> habilita a documentação da linguagem, mesmo se você estiver offline, mostra todos os metodos, os pacotes da linguagem e etc.....

go env -> várias informações do seu ambiente, como o GO esta instalado, qual são as pastas, aonde o GO esta instalado, que esta na variavel GOROOT, outra informação importante é o GOPATH, aonde ele instala o pacote e as dependencias do go, fica ali, varias informacoes da arquitetura e o Sistema operacional que esta utilizando.

go doc cmd/vet -> vai mostrar a documentação desse comando, é uma forma simplificada de identificar código morto, reportar construções suspeitas.

go build -> compilar os programas GO, gera um arquivo, nesse caso com o nome do meu arquivo = comandos, e eu posso executar direto no terminal com o ./comandos;

go run -> compila e já executa -> go run comandos.go

INSTALAR UM PACOTE:

dá um ls dentro de GOPATH
VOU INSTALAR O PACOTE go-sql-driver

go get -u github.com/go-sql-driver/mysql

ls /home/leandro/go/pkg/mod/github.com

*/
