package matematica

import "testing"

const erroPadrao = "Valor esperado $v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) { // método tem que começar como Test e seguir a assinatura de usar um testing
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperado {
		t.Error(erroPadrao, valorEsperado, valor)
	}
}

// o nome do arquivo precisa terminar com "_test"
