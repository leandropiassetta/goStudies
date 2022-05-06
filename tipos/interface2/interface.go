package main

import "fmt"

type esportivo interface {
	ligarTurbo(velocidade int)
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// vou ter que colocar um ponteiro no type ferrari pois irei mexer na variável turboLigado, não posso simplesmente botar ferrari, pq só irei recebe uma cópia dela, e pra modificar tenho que acessar o espaço de memoria onde está ferrari.

// Como eu tenho um método ligarTurbo que não recebe nenhum parâmetro e não retorna nada de fato ferrari agora respeita a interface esportivo, eu posso usar a partir de esportivo uma ferrari...

// esse método agora esta associado ao struct ferrari...
func (f *ferrari) ligarTurbo(velocidade int) {
	f.turboLigado = true
	f.velocidadeAtual = velocidade + f.velocidadeAtual
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo(100)
	fmt.Println(carro1)
	// Na hora que eu atribuir com o tipo esportivo, na hora que for atribuir com a ferrari, tenho que atribuir com endereço..
	var carro2 esportivo = &ferrari{"F40", false, 300}
	carro2.ligarTurbo(400)
	fmt.Println(carro2)
}
