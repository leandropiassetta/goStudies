package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	// Vou fazer um mapeamento como eu quero que esse id vá para o json
	// dados que serão exportados tenho que colocar como maiúsculo
	// identificador que começa com letra maiscula ele tem o significado de ser público.
	// identificador que começa com a letra minúscula ele tem o significado de ser privado.
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	// JSON é um formato textual
	// Método Marshal: retorna dois valores, o primeiro é o texto json e o segundo e o erro
	/*
	   Marshal retorna a codificação JSON de v.

	   Marshal percorre o valor v recursivamente. Se um valor encontrado implementa a interface Marshaler e não é um ponteiro nil, Marshal chama seu método MarshalJSON para produzir JSON. Se nenhum método MarshalJSON estiver presente, mas o valor implementar encoding.TextMarshaler, Marshal chamará seu método MarshalText e codificará o resultado como uma string JSON. A exceção de ponteiro nil não é estritamente necessária, mas imita uma exceção semelhante e necessária no comportamento de UnmarshalJSON.
	*/
	p1Json, _ := json.Marshal(p1) // retorna um slice de bytes
	fmt.Println(string(p1Json))   // converte os bytes em strings

	// json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome":"Caneta", "preco":8.90, "tags":["Papelaria", "Importando"]}`
	// o primeiro argumento um slice de bytes, e vou passar a referencia pro p2, para que ele possa ler o json e setar os valores dentro de produtos
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}
