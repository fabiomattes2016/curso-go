package main

import "fmt"

func main() {
	/*
		A função fmt.Println é uma função variática.
		Uma função varitática trabalha com qualquer número de argumentos.
		Ex.: func Println(a ...interface{}) (n int, err error)
	*/

	numerodebytes, erros := fmt.Println("Hello World", 100)
	fmt.Println(numerodebytes, erros)
}
