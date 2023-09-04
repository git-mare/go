package main

import (
	"consulta-cnpj/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/consultar-cnpj", handlers.ConsultarCNPJHandler)

	port := 8080
	// exemplo de cnpj: 29351447000197
	// exemplo de cnpj com simbolos: 06.947.283/0001-60
	fmt.Printf("Consulta: localhost:%d/consultar-cnpj?cnpj=\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
