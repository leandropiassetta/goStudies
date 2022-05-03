package main

import "fmt"

func main() {
	funcsEsalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsEsalarios["Rafael Filho"] = 1350.0
	delete(funcsEsalarios, "Inexistente")

	for nome, salario := range funcsEsalarios {
		fmt.Println(nome, salario)
	}
}
