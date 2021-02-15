package main

import "fmt"

func main() {

	arrayNomes := [4]string{"Daniel", "Camila", "Vicente"}

	for indice, nome := range arrayNomes {
		fmt.Println(indice)
		fmt.Println(nome)
	}

	for _, nome := range arrayNomes {
		fmt.Println(nome)
	}

	usuario := map[string]string{
		"nome":       "Daniel",
		"sobrenome":  "Aprea",
		"nascimento": "Maio",
	}

	fmt.Println(usuario)

	for _, valor := range usuario {
		fmt.Println(valor)
	}

}
