package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < k; i++ {
		fmt.Print(leftBinarySearch(a, b[i]), " ")
	}
}
func leftBinarySearch(list []int, x int) int {
	cnt := 0
	for i := range list {
		if list[i] == x {
			cnt++
		}
	}
	return cnt
}
