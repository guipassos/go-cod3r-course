package main

import (
	"fmt"
	"time"
)

// Chanel (canal) - é a forma de comunicação entre goroutines
// chanel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // chama a função de maneira independente e a execução do código continua abaixo

	a, b := <-c, <-c // recebendo dados do canal (os dois primeiros valores setados no canal)
	fmt.Println(a, b)

	fmt.Println(<-c) // recebendo o último valor enviado para o canal
}
