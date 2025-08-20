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
	counterSort(array, number)
}
func counterSort(list []int, width int) {
	span := 1001
	counts := make([]int, span)
	for i := 0; i < width; i++ {
		counts[list[i]] += 1
	}
	for i := span - 1; i >= 0; i-- {
		for j := 0; j < counts[i]; j++ {
			fmt.Print(i, " ")
		}
	}
}
