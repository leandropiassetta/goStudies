package main

import (
	"fmt"
)

func main() {
	word := "Ronaldo"

	word = word[:len(word)-1]

	fmt.Println(word + "a")

}
