package operations

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
		result = "o primeiro valor é maior que o segundo"
	} else if a < b {
		result = "o primeiro valor é menor que o segundo"
	} else {
		result = "os numeros são iguais."
	}
	return result
}
