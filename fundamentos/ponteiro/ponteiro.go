package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // criando o ponteiro usando o simbolo * e o tipo int neste caso

	p = &i // pegando o endereço da variavél i (quando mexer na variável i altera o valor de p e vice versa)

	*p++ // desreferenciando  (pegando apenas o valor e incrementando)
	i++

	// Go não tem aritimética de ponteiros, somente nos seus valores após desreferenciar
	// p++

	fmt.Println(p, *p, i, &i)
}
