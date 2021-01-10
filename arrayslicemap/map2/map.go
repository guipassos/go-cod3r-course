package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ // inicia o map já na declaração e permite já atribuir os elementos
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200, // ultimo elemento obrigatoriamente precisa ter virgula ou uma chave encostada no elemento
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // deixa tentar excluir um elemento que não existe (não apresenta erro)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
