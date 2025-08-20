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
		fmt.Println(binarySearch(a, b[j]))
	}
}
func binarySearch(list []int, x int) string {
	left := 0
	right := len(list) - 1
	for left <= right {
		m := left + (right-left)/2
		if list[m] == x {
			return "YES"
		}
		if list[m] < x {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return "NO"
}
