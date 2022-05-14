package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:senha@/goStudies")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // atrasar o fechamento, mas vai fechar a conexao

	tx, _ := db.Begin() // db.Begin, vai inicializar uma transação e vai me dar esse objeto chamado tx, e a partir desse objeto tx que vamos aplicar os comandos sql.
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
