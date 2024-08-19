package main

import "fmt"

func bubbleSort(array []int, len int) []int {
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if array[j] > array[j+1] {
				array[j] += array[j+1]
				array[j+1] = array[j] - array[j+1]
				array[j] = array[j] - array[j+1]
			}
		}
	}
	return array

}

func main() {
	array := [11]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	sortedArray := bubbleSort(array[:], 11)

	fmt.Println(sortedArray)

}
