package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// n sendo a quantidade de números primos que se deseja retornar
func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	defer close(c) // com close estamos dizendo para o compilador que não será enviado mais nenhum dado para o canal e ai o laço for finaliza, se não fizer isso gera um deadlock
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
