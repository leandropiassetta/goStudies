package arquitetura

import (
	"runtime"
	"testing"
)

// É possível dentro de GO, dizer que um determinado codigo fonte vai ser compilado pra arquiterura adm 64, ou um determinado código só vai funcionar no linux e nao no windows, e ai você pode fazer convenções de arquivos ou a partir de comentários você pode dizer que determinado código só vai funcionar numa arquitetura A, ou arquitetura B

func TestDependente(t *testing.T) {
	t.Parallel() // esse teste pode ser executado de forma paralela
	// Arquitetura que estou executando estes testes
	if runtime.GOARCH == "amd64" {
		// pula
		t.Skip("Não funciona em arquitetura amd64")
	}
	//...
	t.Fail()
}
