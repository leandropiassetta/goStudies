// comando para instalar go get -u github.com/go-sql-driver/mysql
package main

// Nesse import não vamos explicitamente usar ele, por isso usa o underline...

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// função exec p/ facilitar minhas execuções
// referencia para sql.DB, conexão que eu abri
// retornar a função sql.Result
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err) // para o programa com o panic(err), passando o erro.
	}

	return result
}

func main() {
	// Abrir uma conexão com o banco de dados, então não vou abrir uma conexão diretamente apontando pro banco de dados que eu quero trabalhar por que esse bd ainda não existe.
	// "mysql" vai ser o banco de dados, qnd eu passso essa string"mysql" ele vai usar indiretamente a dependencia "github.com/go-sql-driver/mysql" que vai ser o driver de acesso ao banco de dados, esse driver ele deve implementar uma série de interfaces definidas pelo GO e ele vai atender quando eu coloco a string "mysql", porque ele vai ser o driver especifico para esse tipo de dados, esse dialeto que é o mysql
	db, err := sql.Open("mysql", "root:senha@/") // @/ vou me conectar com o raiz, não vou passar nenhum bd especificio, pq nesse exemplo quero criar o schema e a tabela.
	if err != nil {
		panic(err)
	}
	// defer -> Aplica essa palavra reservada, e uma sentença de codigo, e essa sentença vai ser prorrogada no final da função.
	// abre uma conexao, e no final da funcao quero que essa conexao seja fechada.
	// não vai ser fechada na linha 34, mas no final da funçao main, logo antes da sua função terminar, essa funcao Close vai ser chamada e essa conexao vai ser fechada.
	defer db.Close()

	exec(db, "create database if not exists goStudies")
	exec(db, "use goStudies") // A partir de todos os comandos daqui pra frente ele vai usar esse comando
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
