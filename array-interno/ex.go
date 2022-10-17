package main

import "fmt"

func main() {
	slice := make([]float32, 10, 15)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // capacidade

}
