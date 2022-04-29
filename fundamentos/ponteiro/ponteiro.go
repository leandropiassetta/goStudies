package main

import "fmt"

// Ponteiro = referência de memória
func main() {
	// Go não tem aritmética de ponteiros...

	i := 1 // A partir dessa variável, a gente ter o acesso a esse endereço dela e compartilhar esse endreço pra colocar esse endereço em um ponteiro

	var p *int = nil // nil -> nulo

	p = &i // Estou dizendo o seguinte pra linguagem GO, pega o endereço dessa variável e me de o valor e atribuo o endereço pro ponteiro p, p guarda o endereço de memoria que aponta pro mesmo local que a variavel i esta apontando;

	// OBS: *p eu tenho o acesso ao valor que o ponteiro aponta, se eu colocar apenas p, eu tenho o endereço que esta armazenado dentro do ponteiro

	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
