package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6) // adiciona elementos no slice, caso seja pequeno o array interno é aumentado automaticamente
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // pega tudo que tem no slice1 e joga para o slice2, como o tamanho do slice2 é apenas 2, será copiado somente as 2 primeiras posições do slice1. O copy não expande o tamanho do array de um slice pra outro. Copy não aceita array, obrigatóriamente tem que ser um slice ou uma string
	fmt.Println(slice2)
}
