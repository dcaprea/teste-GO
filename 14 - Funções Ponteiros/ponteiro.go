package main

import "fmt"

func inverteValor(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	fmt.Println(numero)
	inverteValor(&numero)
	fmt.Println(numero)
}
