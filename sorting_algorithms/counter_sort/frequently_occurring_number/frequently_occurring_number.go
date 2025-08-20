package main

import (
	"fmt"
)

func main() {
	counterSort()
}
func counterSort() {
	span := 10001
	counts := make([]int, span)
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		counts[x]++
	}
	max := counts[0]
	maxNumber := 0
	for i := 0; i < span; i++ {
		if counts[i] >= max {
			maxNumber = i
			max = counts[i]
		}
	}
	fmt.Println(maxNumber)
}
