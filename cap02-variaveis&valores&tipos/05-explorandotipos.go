package main

import (
	"fmt"
)

// declarado x como int
var x int = 10

func main() {
	// se a redeclaração de x por exemplo for de 20.5
	// será retornado um erro, já que o tipo já foi definido como in
	x = 20
	fmt.Printf("Valor de x: %v, %T", x, x)
}
