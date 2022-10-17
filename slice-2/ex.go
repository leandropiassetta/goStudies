package main

import "fmt"

func main() {

	sabores := []string{"Peperoni", "mozzari", "Calabresa", "Portugues", "Frango"}

	sabores[0] = "Azeitona"

	// fatia := sabores[0:2]

	// quando nao sei o tamanho da fatia
	// fatia2 := sabores[2:]

	// fmt.Println(fatia)
	// fmt.Println(fatia2)

	// deletar item em slice

	sabores = append(sabores[:1], sabores[4:]...)

	// make tem custo computacional

	fmt.Println(sabores)

}
