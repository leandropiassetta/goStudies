package main

// Aquela estrutura básica de um pedido de venda, você tem o pedido e os itens do pedido...
type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item // slice de item
}

func main() {
}
