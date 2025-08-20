package main

import (
	"fmt"
)

func main() {
	counterSort()
}
func counterSort() {
	span := 101
	counts := make([]int, span)
	var coinQuantuty int
	fmt.Scan(&coinQuantuty)
	for i := 0; i < coinQuantuty; i++ {
		var coinNominal int
		fmt.Scan(&coinNominal)
		counts[coinNominal]++
	}
	max := 0
	for i := 0; i < len(counts); i++ {
		if counts[i] > max {
			max = counts[i]
		}
	}
	fmt.Println(max)
}
