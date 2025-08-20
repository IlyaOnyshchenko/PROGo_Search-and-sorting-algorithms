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
	array = selectionSort(array)
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i], " ")
	}
}
func selectionSort(list []int) []int {
	for i := len(list) - 1; i > 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			temp := list[maxIndex]
			list[maxIndex] = list[i]
			list[i] = temp
		}
	}
	return list
}
