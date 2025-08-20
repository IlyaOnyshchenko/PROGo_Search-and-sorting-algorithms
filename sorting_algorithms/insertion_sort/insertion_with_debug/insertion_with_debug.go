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
	insertionSort(array)
}
func insertionSort(list []int) {
	var buffer int
	for i := 1; i < len(list); i++ {
		buffer = list[i]
		j := i
		flag := false
		for j > 0 && list[j-1] > buffer {
			list[j] = list[j-1]
			j--
			flag = true
		}
		list[j] = buffer
		if flag == true {
			for i := 0; i < len(list); i++ {
				fmt.Print(list[i], " ")
			}
			fmt.Println()
		}
	}
}
