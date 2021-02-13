package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	fmt.Println("--------------------------------------------------------------")
	variavel1++

	fmt.Println(variavel1, variavel2)

	fmt.Println("--------------------------------------------------------------")

	var variavel3 int = 20
	var variavel4 *int

	variavel4 = &variavel3

	fmt.Println(variavel3, *variavel4)

	variavel3++

	fmt.Println(variavel3, *variavel4)
}
