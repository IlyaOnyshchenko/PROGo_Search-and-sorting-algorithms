package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	array := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	radixSort(array)
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i], " ")
	}
}
func radixSort(list []int) {
	digitCounts := 10
	maxLengthOfNumber := 9
	p := 1
	pocket := make([][]int, digitCounts)
	for i := 0; i < len(pocket); i++ {
		pocket[i] = make([]int, 0)
	}
	for i := 0; i < maxLengthOfNumber; i++ {
		for j := 0; j < len(list); j++ {
			index := list[j] / p % 10
			pocket[index] = append(pocket[index], list[j])
		}
		count := 0
		for j := 0; j < digitCounts; j++ {
			for k := 0; k < len(pocket[j]); k++ {
				list[count] = pocket[j][k]
				count++
			}
			pocket[j] = pocket[j][:0]
		}
		p *= 10
	}
}
