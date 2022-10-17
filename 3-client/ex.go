package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//Contexto com Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	// Preparo a Request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://locahost:8080", nil)

	if err != nil {

		log.Fatalf("Error creating request: %v", err)

	}

	res, err := http.DefaultClient.Do(req) // Aqui executando a minha chamada http

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer res.Body.Close()

	//Stdout aparece no terminal o conteudo do nosso body
	io.Copy(os.Stdout, res.Body) // imprimir na tela o resultado que eu consegui

}
