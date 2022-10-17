package main

import "fmt"

// agrupamento de dados. slice multi-dimensional
func main() {

	sliceSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	lessNumber := sliceSlice[0][0]

	for _, numbers := range sliceSlice {

		for i := 1; i < len(numbers); i++ {
			atual := numbers[i]

			if atual < lessNumber {
				lessNumber = atual
			}
		}
	}
	fmt.Println(lessNumber)

}
