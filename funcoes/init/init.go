package main

import "fmt"

func init() { // função que o compilador chamará antes da função main()
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}

// executar:
// cd funcoes/init
// go run ./
