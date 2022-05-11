package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

//...string , passando parâmetros variáveis...
func titulo(urls ...string) <-chan string {
	// tudo fica a cargo dessa função encapsulado dentro desse padrão generator
	c := make(chan string)

	for _, url := range urls {
		// função anonima a partir de uma goroutine
		go func(url string) {
			// resposta
			// ela não vem como uma string
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			// compilou a exoressao regular que retorna o titulo
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) // invocar a função
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com.br", "https://www.youtube.com")
	fmt.Println("Primeiros", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
