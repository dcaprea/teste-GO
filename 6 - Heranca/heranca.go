package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
}

type aluno struct {
	pessoa
	faculdade string
	curso     string
}

func main() {

	pes := pessoa{"Daniel", "Aprea", 36}

	dadosAluno := aluno{pes, "FMU", "Educação Física"}

	fmt.Println(dadosAluno)
	fmt.Println(dadosAluno.idade)

}
