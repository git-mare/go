package conversor

import (
	"conversor-moeda/formula"
	"fmt"
)

func Start() float64 {
	var optionInput, optionOutput float64
	var inputValue float64

	fmt.Println("Escolha a moeda de origem:")
	fmt.Println("1. Dólar (USD)")
	fmt.Println("2. Euro (EUR)")
	fmt.Println("3. Real (BRL)")

	fmt.Print("Opção: ")
	fmt.Scan(&optionInput)

	fmt.Println("Escolha a moeda de destino:")
	fmt.Println("1. Dólar (USD)")
	fmt.Println("2. Euro (EUR)")
	fmt.Println("3. Real (BRL)")

	fmt.Print("Opção: ")
	fmt.Scan(&optionOutput)

	fmt.Print("Digite o valor a ser convertido: ")
	fmt.Scan(&inputValue)

	result := formula.Conversor(inputValue, optionInput, optionOutput)
	currencyOriginName := formula.GetName(optionInput)
	currencyConversionName := formula.GetName(optionOutput)

	fmt.Printf("%.2f %v equivale a aproximadamente %.2f %v", inputValue, currencyOriginName, result, currencyConversionName)
	return result
}
