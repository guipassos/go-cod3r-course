package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // o tipo dessa variavel é uma interface vazia
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa" // criando uma variavel a partir do tipo dinamico criado
	fmt.Println(coisa2)

	coisa2 = true // como interface é mais genérico pode receber qualquer coisa, semelhante ao object do java
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
