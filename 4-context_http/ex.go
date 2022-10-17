package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	// fazendo request
	// body{} nil
	req, err := http.NewRequest(http.MethodGet, "https://twitter.com", nil)
	if err != nil {
		panic(err)
	}

	// criando um novo contexto
	// queremos colocar um tempo limite para essa resposta...
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))

	// limpeza de recursos, como arquivos abertos, conexoes de rede e conexoes com o banco de dados
	//Quando seu programa for finalizado com esses recursos, é importante fechá-los para evitar exaurir os limites do programa e permitir que outros programas acessem esses recursos.
	defer cancel()
	req = req.WithContext(ctx) // passo o context para dentro da request

	// QUANDO USAR CONTEXT?

	// INTEGRACAO DE SOFTWARES DE TERCEIRO NA SUA APP É ESSENCIAL QUE SE USE ESSA TECNICA
	// DE PASSAR UM TIMEOUT PARA SUA REQUEST, PARA SEU SISTEMA NAO FICAR PRESO NA SUA REQ..

	// iniciar nosso cliente
	res, err := http.DefaultClient.Do(req)

	//validar o erro
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// leitura da resposta
	out, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	//string aqui é um array de bytes
	fmt.Println(string(out))
}
