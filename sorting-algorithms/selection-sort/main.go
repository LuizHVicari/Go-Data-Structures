package main

import "fmt"

func selectionSort(array []int, len int) []int {
	for i := 0; i < len; i++ {
		smaller_value_index := i
		for j := i; j < len; j++ {
			if array[j] < array[smaller_value_index] {
				smaller_value_index = j
			}
		}
		if i != smaller_value_index {
			aux := array[i]
			array[i] = array[smaller_value_index]
			array[smaller_value_index] = aux
		}

	}
	return array

}

func main() {
	array := [11]int{0, 4, 7, 1, 2, 6, 9, 3, 8, 5, 10}

	sortedArray := selectionSort(array[:], 11)

	fmt.Println(sortedArray)

}
