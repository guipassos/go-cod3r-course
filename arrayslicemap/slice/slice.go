package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // isso é um array
	s1 := []int{1, 2, 3}  // isso é um slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array
	s2 := a2[1:3] // pega do indice 1 até o indice 3 sem incluir o indice 3 e o 0
	fmt.Println(a2, s2)

	s3 := a2[:2] //novo slice, mas aponta para o mesmo array interno
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: estrutura com tamanho e um ponteiro interno para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
