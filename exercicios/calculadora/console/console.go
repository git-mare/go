package console

import (
	"calculadora/operations"
	"fmt"
	"os"
	"reflect"
)

func Answers() int {
	var num1, num2, result int
	var op string

	fmt.Print("Digite um numero: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Insira apenas valores inteiros.")
		os.Exit(0)
	}

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Digite outro numero: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Insira apenas valores inteiros.")
		os.Exit(0)
	}

	switch op {
	case "+":
		if reflect.TypeOf(num1).Kind() == reflect.Int && reflect.TypeOf(num2).Kind() == reflect.Int {
			result = operations.Add(num1, num2)
		}
	case "-":
		if reflect.TypeOf(num1).Kind() == reflect.Int && reflect.TypeOf(num2).Kind() == reflect.Int {
			result = operations.Sub(num1, num2)
		}
	case "*":
		if reflect.TypeOf(num1).Kind() == reflect.Int && reflect.TypeOf(num2).Kind() == reflect.Int {
			result = operations.Mult(num1, num2)
		}
	case "/":
		if reflect.TypeOf(num1).Kind() == reflect.Int && reflect.TypeOf(num2).Kind() == reflect.Int {
			result = operations.Div(num1, num2)
		}
	default:
		fmt.Println("Operador inválido ou valores não inteiros inseridos.")
		return -1
	}

	fmt.Println("Resultado da operação: ", result)
	comp := operations.Comp(num1, num2)
	fmt.Println("Comparação: ", comp)

	return result
}
