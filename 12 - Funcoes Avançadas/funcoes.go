package main

import "fmt"

func calculosMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func somaNnumeros(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	soma, subtracao := calculosMatematico(10, 20)
	fmt.Println(soma, subtracao)
	somando := somaNnumeros(1, 2, 3, 100, 23, 200)
	fmt.Println(somando)

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Teste de texto")

	novotestetexto := func(texto2 string) string {
		return fmt.Sprintf("Retorno %s", texto2)
	}("Teste de texto 2")
	fmt.Println(novotestetexto)
}
