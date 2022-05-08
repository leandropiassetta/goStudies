package main

import "math"

// Iniciando com a letra maiúscula é PÚBLICO (visivel dentro e fora do pacote)!

// Dentro de um pacote, eu posso ter vários arquivos, não importa muito quantos arquivos voce colocar, no final o arquivo em si não será relevante porque isso será compilado numa unica estrutura, que é a estrutura do pacote... não pode ter funcões duplicadas entre vários arquivos dentro de um mesmo pacote, uma função tem um identificador único, um nome único dentro de um determinado pacote, no GO, voce nao tem visibilidade privada dentro de um arquivo, e sim dentro de um pacote..

// Iniciando com letra minúscula é PRIVADO (visivel dentro do pacote)!, não dentro do arquivo e sim dentro do pacote.. quando é compilado esses arquivos deixam de existir e passam a ter apenas o pacote, com todas as estruturas que eu defini la dentro..

// Por exemplo...
// Ponto - gerará algo público..
// ponto ou _Ponto - gerará algo privado..

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	// Público -> tem que ter comentários
	x float64
	y float64
}

// to fazendo retorno nomeado
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return // retorno limpo
}

// Distancia é repsonsavel por calcular a ditância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	// Raiz quadrada da soma dos quadrados dos catetos
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
