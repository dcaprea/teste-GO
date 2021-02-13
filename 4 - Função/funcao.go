package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2

}

var fun2 = func(txt string) string {
	fmt.Println(txt)
	return txt
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func calculosMatematicos2(n1, n2 int8) (int8, int8) {
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return multiplicacao, divisao
}

func main() {
	soma := somar(20, 20)
	fmt.Println(soma)

	var fun = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fmt.Println(fun("Daniel Aprea"))
	fmt.Println(fun2("Camila Aprea"))

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtracao)

	resultadomultiplicacao, _ := calculosMatematicos2(40, 20)
	fmt.Println(resultadomultiplicacao)
}
