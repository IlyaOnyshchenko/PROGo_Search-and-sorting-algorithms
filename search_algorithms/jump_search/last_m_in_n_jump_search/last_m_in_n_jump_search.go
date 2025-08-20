package main

import (
	"fmt"
	"math"
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
	for i := 0; i < k; i++ {
		fmt.Println(jumpSearch(a, b[i]))
	}
}
func jumpSearch(list []int, x int) int {
	b := int(math.Sqrt(float64(len(list))))
	start := 0
	end := b - 1
	for list[end] <= x && end < len(list)-1 && list[end+1] <= x {
		start = int(math.Min(float64(len(list)-1), float64(start+b)))
		end = int(math.Min(float64(len(list)-1), float64(end+b)))
	}
	if x > list[end] {
		return 0
	}
	for i := end; i >= start; i-- {
		if list[i] == x {
			return i + 1
		}
	}
	return 0
}
