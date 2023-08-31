package formula

import (
	"fmt"
	"os"
)

func Conversor(value, currencyOrigin, currencyConversion float64) float64 {
	exchangeRateBRL := 4.89
	exchangeRateEUR := 0.92

	if currencyConversion == 1 { // Converter para Dolar
		switch currencyOrigin {
		case 2: // EUR para USD
			dolar := value / exchangeRateEUR
			return dolar

		case 3: // BRL para USD
			dolar := value / exchangeRateBRL
			return dolar

		default:
			fmt.Println("Entrada inválida")
			os.Exit(0)
		}
	} else if currencyConversion == 2 { // Converter para Euro
		switch currencyOrigin {
		case 1: // USD para EUR
			eur := value * exchangeRateEUR
			return eur

		case 3: // BRL para EUR
			dolar := value / exchangeRateBRL
			eur := dolar * exchangeRateEUR
			return eur

		default:
			fmt.Println("Entrada inválida")
			os.Exit(0)
		}
	} else if currencyConversion == 3 { // Converter para Real
		switch currencyOrigin {
		case 1: // USD para BRL
			dolar := value * exchangeRateBRL
			return dolar

		case 2: // EUR para BRL
			dolar := value / exchangeRateEUR
			brl := dolar * exchangeRateBRL
			return brl

		default:
			fmt.Println("Entrada inválida")
			os.Exit(0)
		}
	}

	fmt.Println("Entrada inválida")
	os.Exit(0)
	return -1
}

func GetName(currency float64) string {
	var currencyName string
	switch currency {
	case 1:
		currencyName = "USD"
	case 2:
		currencyName = "EUR"
	case 3:
		currencyName = "BRL"
	default:
		fmt.Println("Entrada inválida")
		os.Exit(0)
	}
	return currencyName
}
