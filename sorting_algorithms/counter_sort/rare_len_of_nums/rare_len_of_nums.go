package main

import (
	"fmt"
)

func main() {
	counterSort()
}
func counterSort() {
	var span int = 100000001
	var cnt int
	counts := make([]int, span)
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		counts[x]++
	}
	maxNumberLength := 1000
	minLength := 9
	for i := span - 1; i >= 0; i-- {
		if (decLength(i) <= minLength) && (counts[i] <= maxNumberLength) && (counts[i] != 0) {
			maxNumberLength = counts[i]
			minLength = decLength(i)
		}
	}
	for i := intPow(10, minLength-1); i < intPow(10, minLength); i++ {
		if counts[i] > 0 {
			cnt++
		}
	}
	fmt.Println(minLength, cnt)
}
func decLength(num int) int {
	var cnt int
	for num > 0 {
		num /= 10
		cnt++
	}
	return cnt
}
func intPow(osn int, stepen int) int {
	var rez int = 1
	for i := 0; i < stepen; i++ {
		rez *= osn
	}
	return rez
}
