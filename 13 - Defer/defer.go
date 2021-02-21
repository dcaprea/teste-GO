package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada, resultado será retornado")
	fmt.Println("Aluno está aprovado")
	resultado := (n1 + n2) / 2
	if resultado >= 6 {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(alunoAprovado(7, 8))
}
