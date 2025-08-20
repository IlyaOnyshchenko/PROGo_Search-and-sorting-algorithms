package main

import (
	"fmt"
)

func main() {
	counterSort()
}
func counterSort() {
	var colomnQuantity int
	fmt.Scan(&colomnQuantity)
	counts := make([]int, colomnQuantity)
	var squareQuantuty int
	fmt.Scan(&squareQuantuty)
	for i := 0; i < squareQuantuty; i++ {
		var colomnNumber int
		fmt.Scan(&colomnNumber)
		counts[colomnNumber-1]++
	}
	min := squareQuantuty
	for i := 0; i < len(counts); i++ {
		if counts[i] < min {
			min = counts[i]
		}
	}
	fmt.Println(min)
}
