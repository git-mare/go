package main

import (
	"calculadora/operations"
	"fmt"
)

func main() {
	op := operations.Add(1, 2)
	comp := operations.Comp(1, 2)
	fmt.Printf("%v\n%v", op, comp)
}
