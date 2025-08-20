package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	b := make([]int, k)
	for i := range n {
		fmt.Scan(&a[i])
	}
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < k; i++ {
		fmt.Println(findClosest(a, b[i]))
	}
}
func findClosest(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	closest := arr[0]
	for low <= high {
		mid := low + (high-low)/2
		if abs(arr[mid]-target) < abs(closest-target) {
			closest = arr[mid]
		}
		if abs(arr[mid]-target) == abs(closest-target) {
			closest = min(closest, arr[mid])
		}
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			return arr[mid]
		}
	}
	return closest
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
