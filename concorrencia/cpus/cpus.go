package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Esse método mostra quantos processadores fisícos tem na sua máquina
	// Mesmo com 4 processadores a concorrência pode ocorrer por que ela não depende de múltiplos processadores fisícos, ele habilita você poder ter paralelismo ou nao;...
	fmt.Println(runtime.NumCPU())
}
