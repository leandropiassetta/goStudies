package matematica

// Testes unitários, testar uma unidade, por exemplo uma função, quanto mais você isola essa função pra testar ela de uma forma isolada do mundo exsterior, você trata essa funçao de forma unica..

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular o que você já sabe... :)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	fmt.Println(media)
	// strconv -> converte uma string em um float
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
