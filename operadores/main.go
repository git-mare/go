package main

import "fmt"

var n1 = 5 // declaração fora do codeblock

func main() {

	fmt.Println("valor declarado fora do codeblock", n1)

	n1 := 10 // declaração dentro do codeblock
	fmt.Println("valor declarado dentro do codeblock:", n1)

	n1 = 15
	fmt.Println("valor atribuido", n1)

	n2 := 10 == 10 // operação de comparação
	fmt.Println(n2)

}
