package main

import "fmt"

func main() {
  x := "oi "
  y := "bom dia"
  z := fmt.Sprint(x, y) // fmt.Sprint() armazena em string as variáveis x e y

  fmt.Println(z) // Imprime a variável z que está armazenando a Sprint.
}
