package main

import (
	"fmt"
)

func main() {
	var n, minEl, maxEl int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	minEl = a[0]
	maxEl = a[0]
	for i := 1; i < n; i++ {
		if a[i] < minEl {
			minEl = a[i]
		} else if a[i] > maxEl {
			maxEl = a[i]
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == minEl {
			a[i] = maxEl
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}
