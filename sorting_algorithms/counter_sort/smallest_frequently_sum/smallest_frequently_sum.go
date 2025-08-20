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
	maxEncounter := 0
	minSumm := 36
	for i := span - 1; i >= 0; i-- {
		if (decSumm(i) <= minSumm) && (counts[i] >= maxEncounter) && (counts[i] != 0) {
			maxEncounter = counts[i]
			minSumm = decSumm(i)
		}
	}
	fmt.Println(minSumm)
}
func decSumm(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
