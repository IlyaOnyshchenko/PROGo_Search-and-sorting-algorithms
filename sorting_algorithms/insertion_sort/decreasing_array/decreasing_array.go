package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	array = insertionSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
}
func insertionSort(list []int) []int {
	var buffer int
	for i := 1; i < len(list); i++ {
		buffer = list[i]
		j := i
		for j > 0 && list[j-1] < buffer {
			list[j] = list[j-1]
			j--
		}
		list[j] = buffer
	}
	return list
}
