package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log.Println("Iniciou minha request")

	defer log.Println("Finalizou minha request")

	select {

	case <-time.After(time.Second * 5):
		fmt.Fprintln(w, "Requisição processada com sucesso!")
		w.Write([]byte("Requisição processada com sucesso!")) // client recebe essa resposta

	case <-ctx.Done():
		log.Println("Request cancelada")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}

}
