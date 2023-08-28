package main

import (
	"fmt" // biblioteca importada da própria linguagem GO

	"exemplo-package/math" // diretório importado da pasta do projeto
)

func main() {
	result := math.Add(1, 2)

	fmt.Println("Soma: ", result)
}
