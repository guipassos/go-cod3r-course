package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}

// o canal é um bloqueio e se você não tiver mais valor no canal ele ficará bloqueado pra sempre gerando o deadlock
