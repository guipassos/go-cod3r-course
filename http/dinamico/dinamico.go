package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") // parece uma data/hora mas na verdade cada número tem um significado semelhante ao padrão dd/mm/yyyy. 02 => dia, 01 => mes, 2006 => ano, 03 => hora, 04 => minuto, 05 => seg
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)  // usando a função Fprintf para escrever no writer
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta) // url mapeada
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
