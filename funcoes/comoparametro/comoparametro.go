package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int { // variável como funcao tem que se informar quais os tipos de dados esperados dessa função (a assinatura da função, o que ele espera com parametro e qual o seu retorno)
	return funcao(p1, p2) // retorna a funcao passada usando o p1 e p2
}

func main() {
	resultado := exec(multiplicacao, 3, 4) // passa no primeiro parametro a função que será executada e no segundo e terceiro parametros os valores
	fmt.Println(resultado)
}
