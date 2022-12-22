package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 5

var contador int

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\n Total do contador:\t", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}
