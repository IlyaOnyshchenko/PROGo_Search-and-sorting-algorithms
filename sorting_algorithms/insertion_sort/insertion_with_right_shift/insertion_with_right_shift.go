package main

import (
	"fmt"
)

func main() {
	var number int
	var target, index int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Scan(&target, &index)
	array = insertionSort(array, target, index)
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
}
func insertionSort(list []int, element int, position int) []int {
	list = append(list, element)
	for j := len(list) - 1; j > position-1; j-- {
		list[j], list[j-1] = list[j-1], list[j]
	}
	return list
}
