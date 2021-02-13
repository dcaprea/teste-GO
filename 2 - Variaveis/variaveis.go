package main

import "fmt"

func main() {
	var variavel1 string = "Teste"
	variavel2 := "Teste 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Teste 3"
		variavel4        = "Teste 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

}
