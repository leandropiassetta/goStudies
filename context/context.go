package context

import (
	"context"
	"fmt"
	"time"
)

// Context ->> SQL eu quero rodar uma query que nao demore muido, definindo um tempo limite pra ele, bota um timeout no context e resolve o problema.

// Podemos com os deadlines, definir tempo máximo para o código terminar de ser executado, ou então, ser cancelado por timeout do contexto, ou então, ser cancelado por timeout do contexto.

func main() {
	ctx, cancel := context.WithCancel(
		context.Background(), // inicializando context
	)

	go printUntilCancel(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {

		case <-ctx.Done():
			fmt.Println("Cancel signal received, exiting")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Printing until cancel, number = %d \n", count)
			count++
		}
	}

}
