package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)
	fmt.Println(jumpSearch(a, k))
}
func jumpSearch(list []int, x int) int {
	b := int((math.Sqrt(float64(len(list)))))
	start := 0
	end := b - 1
	cnt := 0
	for list[end] <= x && end < len(list)-1 && list[end+1] <= x {
		for i := end; i >= start; i-- {
			if list[i] == x {
				cnt++
			}
		}
		start = min(len(list)-1, start+b)
		end = min(len(list)-1, end+b)
	}
	if x > list[end] {
		return 0
	}
	for i := end; i >= start; i-- {
		if list[i] == x {
			cnt++
		}
	}
	return cnt
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
