package matematica

import (
	"fmt"
	"strconv"
)

// Media é reponsável por calcular o que você já sabe... :)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64) // se trocar o 2f para 3f por exemplo gerará uma falha que será pega no teste
	return mediaArredondada
}
