package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // convertendo o html para string e aplicando a expressão regular do regex e retornando para o canal
		}(url) // passando a url de fora para dentro do go func
	}
	return c // a função irá retornar o canal antes mesmo do go func estar pronto
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2) // o primeiro titulo que chegar no t1 e no t2 cairão aqui
	fmt.Println("Segundos:", <-t1, "|", <-t2) // o segundo titulo que chegar no t1 e no t2 cairão aqui
}
