package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e Slices")

	var array1 [5]string
	array1[0] = "Daniel"
	fmt.Println(array1)

	array2 := [5]string{"Daniel", "Camila", "Vicente", "Irlene", "Antonio"}
	fmt.Println(array2)

	array2[0] = "Benjamim"

	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 60)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	fmt.Println("----------------------------------------------------------------------")

	slice3 := make([]int, 5, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
