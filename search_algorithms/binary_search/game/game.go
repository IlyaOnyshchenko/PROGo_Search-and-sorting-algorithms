package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(binarySearch(n))
}
func binarySearch(x int) int {
	cnt := 0
	for x > 0 && x > 1 {
		x = x/2 + x%2
		cnt++
	}
	return cnt
}
