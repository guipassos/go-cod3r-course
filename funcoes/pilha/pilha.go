package main

import "runtime/debug"

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}

// é comum que linguagens exibam a pilha de execução ao ocorrer algum erro, neste caso não ocorreu erro, mas fi usado uma rotina pra exibir a pilha de execução.
