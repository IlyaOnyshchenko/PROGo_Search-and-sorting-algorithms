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
	min := 1001
	max := -1001
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}
	counts := make([]int, max+intAbs(min)+1)
	for i := 0; i < width; i++ {
		counts[list[i]+intAbs(min)] += 1
	}
	for i := 0; i < (max+intAbs(min))+1; i++ {
		for j := 0; j < counts[i]; j++ {
			if counts[i] != 0 {
				fmt.Print(i-intAbs(min), " ")
			}
		}
	}
}
func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
