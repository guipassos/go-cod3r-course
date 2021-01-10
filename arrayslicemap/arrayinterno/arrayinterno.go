package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // slice criado com tamanho 10 e array interno com capacidade 20
	s2 := append(s1, 1, 2, 3) // slice adicionando mais 3 valores no final
	fmt.Println(s1, s2)

	s1[0] = 7 // embora tenha sido alterando o valor da posição 0 pra 7 no slice 1, o mesmo valor foi atribuido na mesma posição do slice 2. Isso prova que ambos os slices referenciam um mesmo array internamente
	fmt.Println(s1, s2)
}
