package main

import (
	"context"
	"fmt"
	"time"
)

// contexto trabalham como uma árvore, um contexto dentro do outro, um contexto pai,varios filhos e os filhos podem ter filhos, quando um emite um sinal, ele pode sair e parando todos os outros contextos, para que mantenhamos o controle da aplicação, a nossa app sempre vai ter um contexto principal, um contexto branco chamado de context background

func main() {

	// contexto em branco, contexto raiz, o contexto pai
	ctx := context.Background()

	//WithCancel -> Vai retornar uma funcao de cancelamento, quando executar a funcao de cancelamento, todo mundo que tiver aquele contexto vai receber um sinal que aquele contexto ja era, por boa pratica sempre se cancela o contexto no final da execução.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // executa por ultimo

	go func() {
		time.Sleep(time.Second * 6)
		cancel()
	}()

	bookHotel(ctx)

}

//OBS: tenho duas API, uma api A e uma api B, que fazem a reserva do mesmo quarto, mas eu quero ser o mais rapido, manda a requisicao pras duas API, e a que terminar antes, reserva e da um cancel, e a outra API para instantaneamente pois a outra API conseguiu antes, isso tudo para nao gastar mais recursos do processador...

func bookHotel(ctx context.Context) {
	// iniciar um processo de reserva no hotel, se demorar mto pra realizar esse processo de reserva, ele vai abortar e depois desolva para o usuario o tempo limite foi excedido, tente novamente mais tarde.

	// fica aguardando alguma condiçao da certo, alguma condição apita e ele executa, o dá certo, sao channels.
	select {

	// Se o contexto receber um sinal, ele ja foi cancelado
	case <-ctx.Done():
		fmt.Println("tempo excedido para bookar o quarto")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}

}
