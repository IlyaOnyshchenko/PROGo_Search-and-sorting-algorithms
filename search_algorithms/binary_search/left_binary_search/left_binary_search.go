package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	b := make([]int, k)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
	}
	for j := 0; j < k; j++ {
		fmt.Println(leftBinarySearch(a, b[j]))
	}
}
func leftBinarySearch(list []int, x int) int {
	left := -1
	right := len(list) - 1
	for left+1 < right {
		m := left + (right-left)/2
		if list[m] < x {
			left = m
		} else {
			right = m
		}
	}
	if list[right] == x {
		return right + 1
	}
	return 0
}
