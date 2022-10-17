package main

import (
	"fmt"
	"sort"
)

func somaSlice(slice []int) {

	total := 0

	for _, value := range slice {
		total += value
	}

	fmt.Println(total)
}

func maiorNumeroSlice(sliceInt []int) {

	sort.SliceStable(sliceInt, func(i, j int) bool {
		return sliceInt[i] > sliceInt[j]
	})

	fmt.Println(sliceInt[0])
}

func menorNumberOtherWay(slice []int) {

	if len(slice) < 1 {
		panic("não tem elemento neste slice")
	}

	lessNumber := slice[0]

	for i := 0; i < len(slice); i++ {

		atualNumber := slice[i]

		if atualNumber < lessNumber {
			lessNumber = atualNumber
		}
	}
	fmt.Println(lessNumber)
}

func main() {
	array := [5]int{9, 2, 3, 4, 5}

	fmt.Println(array)

	slice := []int{3, 3, 3, 3, 2}

	somaSlice(slice)
	maiorNumeroSlice(slice)
	menorNumberOtherWay(slice)

	// array2 := append(array, 6)
	slice = append(slice, 6)

	fmt.Println(slice)
	// fmt.Println(array2)

	sliceStr := []string{"banana", "maçã", "morango"}

	for indice, valor := range sliceStr {
		fmt.Println("No indíce", indice, "temos o valor: ", valor)
	}

	sliceStr[3] = "melancia"

}
