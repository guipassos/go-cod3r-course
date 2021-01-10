package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // daqui em diante só vai executar quando os anteriores forem consumidos (devido ao buffer de 3 posições)
	fmt.Println("Executou!")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // com buffer com 3 posições
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
