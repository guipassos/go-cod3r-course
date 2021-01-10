package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um slice com 10 elementos
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // internamente a capacidade desse slice passa a ser 20, ou seja é tratado internamente um array de 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // prrenchendo toda a capacidade do slice de 20 posições, começando a partir da décima posição
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // não da erro pois o slice cresce automaticamente quando se atinge a capacidade interna, porém cria internamente um novo array quando o antigo estrapola
	fmt.Println(s, len(s), cap(s))
}
