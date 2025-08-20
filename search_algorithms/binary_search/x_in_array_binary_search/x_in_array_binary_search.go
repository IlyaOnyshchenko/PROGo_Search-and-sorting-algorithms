package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)
	fmt.Println(leftBinarySearch(a, k))
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
