package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	fmt.Print("Insira seu peso (KG): ")
	fmt.Scan(&weight)

	fmt.Print("Insira sua altura (M): ")
	fmt.Scan(&height)

	imc := weight / (height * height)

	if imc < 18.5 {
		fmt.Printf("Abaixo do peso.")
	} else if 18.5 <= imc && imc < 24.9 {
		fmt.Printf("Peso normal.")
	} else if 29.9 <= imc {
		fmt.Printf("Obesidade.")
	}

}
