package main

import "fmt"

func main() {

	// primeiro exemplo
	a := "Go"
	fmt.Printf("Bem vindo a linguagem %v.", a) // O parametro v está imprimindo a variável a ao lado do texto

	// segundo exemplo
	b := "\noi "
	c := "bom dia"
	d := fmt.Sprint(b, c) // fmt.Sprint() armazena em string as variáveis x e y
	fmt.Println(d)        // Imprime a variável z que está armazenando a Sprint.
}
