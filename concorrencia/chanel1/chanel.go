package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita) (enviando 1 para o canal)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2

	fmt.Println(<-ch)
}

// o chanel pode servir para fazer uma sincronismo entre 2 rotinas
