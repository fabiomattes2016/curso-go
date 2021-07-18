/*
	O operador curto de declaração -> :=
	parece uma marmota (gopher) ou o PUNISHER da Marvel
	e é usado para declaração da variável
	Não confundir com operador de atribuição -> =
*/

package main

import "fmt"

// declaração de variável a nível de package
var y = "olá bom dia"

func main() {
	// declaração de varável a nível de escopo
	x := 10.0

	fmt.Printf("Valor de x: %v, Tipo de x: %T\n", x, x)
	fmt.Printf("Valor de y: %v, Tipo de y: %T\n", y, y)
}
