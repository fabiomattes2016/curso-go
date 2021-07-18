/*
Analogia acerca de inicialização vs. atribuição de valores em variáves: Caixas Postais

-- Declaração --
Vou ao correio e solicito uma caixa postal.
Depois de passar pela burocracia e pagamentos, eu recebo um endereço de correspondência.

-- Inicialização --
A primeira correspondência a ser recebida na caixa postal.

-- Atribuição --
Trocar a correspondência da caixa postal por outra.
*/

package main

import (
	"fmt"
)

var x int // <- declaração, reserva de endereço na memória

var a int
var b string
var c float64
var d bool

func main() {
	x = 10 // <- inicialização & atribuição
	fmt.Printf("%v, %T\n", x, x)

	x = 15 // <- atribuição de valor
	fmt.Printf("%v, %T\n", x, x)

	numerodebytes, _ := fmt.Println(x)
	fmt.Printf("Número de bytes usados em x: %v\n\n", numerodebytes)

	fmt.Printf("A: %v, %T\n", a, a)
	fmt.Printf("B: %v, %T\n", b, b)
	fmt.Printf("C: %v, %T\n", c, c)
	fmt.Printf("D: %v, %T\n", d, d)
}
