package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{ // cria um primeiro map cuja chave será uma string, o valor deste primeiro map será outro map (segundo map), cuja chave deste segundo map também é uma string e o valor deste segundo map sendo um float64
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcsPorLetra, "P")
	for letra, funcionarios := range funcsPorLetra {
		fmt.Println(letra, funcionarios)
	}
}
