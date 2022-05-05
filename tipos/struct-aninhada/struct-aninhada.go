package main

import "fmt"

// Aquela estrutura básica de um pedido de venda, você tem o pedido e os itens do pedido...
type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	// struct aninhado
	itens []item // slice de item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, quantidade: 2, preco: 12.10}, // seguir esse tipo de padrão, mais fácil de identificar o que é cada valor no meu slice.
			{2, 1, 23.49},
			{3, 100, 3.13},
		},
	}
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
