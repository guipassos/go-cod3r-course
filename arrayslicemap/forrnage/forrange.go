package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // array com números fixos, onde a contagem / tamanho é feita pelo compilador quando o array é criado; se retirar os tres pontos vira um slice (não é o foco desta aula)

	for i, numero := range numeros { // o primeiro elemento é o key, o segundo o value
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // underline ignora o indice permitindo acessar apenas o elemento de cada posição
		fmt.Println(num)
	}
}
