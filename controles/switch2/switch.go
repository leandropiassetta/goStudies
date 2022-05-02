package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()     // hora atual
	switch /* true */ { // switch true --> Se não colocar o switch vinculado a algum caso, ele vai procurar algum case, que seja verdadeiro, ai irei botar alguns cases com verdadeiros ou falso e o primeiro verdadeiro que ele encontrar, ele vai executar..
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}
}
