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
	for i := 0; i < n; i++ {
		if m == a[i] {
			fmt.Print(i+1, " ")
		}
	}
}
