package operations

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}

func Mult(a, b int) int {
	return a * b
}

func Comp(a, b int) string {
	var result string

	if a > b {
		result = fmt.Sprintf("%v é maior que %v", a, b)
	} else if a < b {
		result = fmt.Sprintf("%v é menor que %v", a, b)
	} else {
		result = fmt.Sprintf("Os numeros são iguais. (%v)", a)
	}

	return result
}
