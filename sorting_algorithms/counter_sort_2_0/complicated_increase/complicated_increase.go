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
	min := 2000000001
	max := -2000000001
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}
	var counts []int
	if min > 0 {
		counts = make([]int, max-min+1)
		for i := 0; i < width; i++ {
			counts[list[i]-min] += 1
		}
	} else if min < 0 {
		counts = make([]int, max-min+1)
		for i := 0; i < width; i++ {
			counts[list[i]+intAbs(min)] += 1
		}
	}
	if min < 0 {
		for i := (max + intAbs(min)) - 107; i < (max+intAbs(min))+1; i++ {
			for j := 0; j < counts[i]; j++ {
				if counts[i] != 0 {
					fmt.Print(i-intAbs(min), " ")
				}
			}
		}
	} else if min > 0 {
		for i := 0; i < (max-min)+1; i++ {
			for j := 0; j < counts[i]; j++ {
				if counts[i] != 0 {
					fmt.Print(i+min, " ")
				}
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
