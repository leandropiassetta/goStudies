package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:*Tigre91@/goStudies") // obs: conecto com o banco que eu criei antes
	if err != nil {
		panic(err)
	}

	defer db.Close() // para garantir que antes de sair da funcao main, o  Close vai ser chamado para fechar a conexão
	// statement = stmt

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria jose") // quando eu for executar esses statement, ai eu passo o valor que eu quero que ele substitua nessa interrogação;
	stmt.Exec("João")       // função Exec retorna uma resposta, inclusive nessa resposta, vem o id do item que acabou de ser inserido no banco de dados.

	res, _ := stmt.Exec("joao")

	id, _ := res.LastInsertId() // pego o ultimo id inserido
	fmt.Println(id)

	linhas, _ := res.RowsAffected() // pego a quantidade de linhas que foram afetadas com esse comando.
	fmt.Println(linhas)
}
