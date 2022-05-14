package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// para mapear nosso usuário
type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:senha@/goStudies")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db.Query(): consulta e retorna linhas
	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
	defer rows.Close()

	// Para ele pegar cada um dos elementos enquanto tiver esses elementos.
	for rows.Next() {
		// chamamos o método Scan, para scanear o result set e aplicar isso numa struct
		var u usuario
		// ai ele está mapeando o retorno dessa consulta, para o valor desse usuario que defini ali em cima
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
