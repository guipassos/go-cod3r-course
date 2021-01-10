package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado pois desta maneira não irá concatenar o número e sim pegar o valor correspondente na tabela ascii...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // função que retorna 2 valores, o primeiro é o valor e o segundo é um erro, neste caso o simbolo _ esta ignorando esse erro que poderá retornar dependendo a conversão
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("F")

	if b {
		fmt.Println("Verdadeiro")
	}
}
