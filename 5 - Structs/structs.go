package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	fmt.Println("Estudando Struct")

	enderecoExemplo := endereco{"Rua Teste", 200}

	var u usuario
	fmt.Println(u)
	u.nome = "Daniel"
	u.idade = 36
	u.endereco = enderecoExemplo

	fmt.Println(u)

	u2 := usuario{"Camila", 33, enderecoExemplo}
	fmt.Println(u2)
}
