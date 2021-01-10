package main

import "fmt"

func media(numeros ...float64) float64 { // permite passar vários valores apenas setando 3 pontos, tornando o parametro variável
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
