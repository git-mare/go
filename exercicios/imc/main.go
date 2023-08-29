package main

import (
	"fmt"
)

func main() {
	var weight, height float64
	var category string

	fmt.Print("Insira seu peso (KG): ")
	fmt.Scan(&weight)

	fmt.Print("Insira sua altura (M): ")
	fmt.Scan(&height)

	imc := weight / (height * height)

	if imc < 18.5 {
		category = fmt.Sprintf("Abaixo do peso.")
	} else if 18.5 <= imc && imc < 24.9 {
		category = fmt.Sprintf("Peso normal.")
	} else if 29.9 <= imc {
		category = fmt.Sprintf("Obesidade.")
	}

	fmt.Printf("Seu IMC é %.1f\n", imc)
	fmt.Println("Categoria: ", category)

}
