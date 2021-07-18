package main

import (
	"fmt"
)

var z = 10

func main() {
	y := 10 + z
	qualquercoisa(y)
}

func qualquercoisa(x int) {
	fmt.Println(x)
}
