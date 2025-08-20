package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&m)
	if linearSearch(a, m) != 0 {
		fmt.Println(linearSearch(a, m))
	}
}
func linearSearch(list []int, key int) int {
	for index, value := range list {
		if key == value {
			return index + 1
		}
	}
	return 0
}
