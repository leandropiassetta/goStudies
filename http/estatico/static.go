package main

import (
	"log"
	"net/http"
)

func main() {
	// como parametro para o FileServer o diretorio a partir da qual ele vai ler
	fs := http.FileServer(http.Dir("public"))
	// Handle() -> digo que o "/" automaticamente vai passar para o fs
	// Quando chegar uma requisição na raíz da minha aplicação, automaticamente passe para esse Handle, que é o fs que eu criei na linha 7
	http.Handle("/", fs)

	log.Println("Executando....")                // pra dizer que o servidor esta executando
	log.Fatal(http.ListenAndServe(":3000", nil)) // Pra ele realmente startar o servidor e ficar escutando na porta 3000
}

// pra executar corretamente, para ele entender de que diretorio estou executando, tem que ser feito pelo terminal, usando os comandos
// cd http/static
// go run static.go
