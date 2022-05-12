package matematica

import "testing"

//%v -> value
const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// *testing.T -> estrutura auxiliar(helper), um ponteiro pra essa estrutura
// testando no terminal -> cd pastaDoTeste go test ou go test ./... -> olha a diretorio atual, mas tb as pastas de baixa desse diretorio e executa os testes que estão ali.

// OBS: seguir a convenção os métodos de testes começam sempre com Test e o arquivo termina com o _test.go
func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
