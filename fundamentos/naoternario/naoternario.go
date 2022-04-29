package main

import "fmt"

// Não tem operador ternário

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	// return nota >- 6 ? "Aprovado" : "Reprovado" -> Operador Ternário
}

func main() {
	fmt.Println(obterResultado(6.2))
}

// se quiser operador ternario usa o else
