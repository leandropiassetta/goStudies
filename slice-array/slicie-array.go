package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAY
	var array1 [5]int
	array1[0] = 10

	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}

	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array1, array2, array3)

	//SLICE

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}

	slice = append(slice, 18)

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

}
