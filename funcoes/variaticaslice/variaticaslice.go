package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"} //slice
	imprimirAprovados(aprovados...)                                // ao usar os "..." a linguagem já irá explodir os valores como parâmetros da função imprimirAprovados
}
