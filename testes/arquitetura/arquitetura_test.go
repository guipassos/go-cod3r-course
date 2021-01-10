package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // otimiza o tempo de execução dos testes quando houver muitos testes
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64") // pula o teste
	}
	// ...
	t.Fail()
}
