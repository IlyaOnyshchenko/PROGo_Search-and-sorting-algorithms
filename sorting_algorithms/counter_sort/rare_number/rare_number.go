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
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		counts[x]++
	}
	min := 10000
	minNumber := 101
	for i := span - 1; i >= 0; i-- {
		if (counts[i] <= min) && (counts[i] != 0) {
			minNumber = i
			min = counts[i]
		}
	}
	fmt.Println(minNumber)
}
