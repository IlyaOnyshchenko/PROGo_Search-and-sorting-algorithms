package main

import (
	"fmt"
	"math"
)

func main() {
	var m, razn float64
	var n, elem int
	fmt.Scan(&n)
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&m)
	razn = math.Abs(m - a[0])
	elem = int(a[0])
	for i := 1; i < n; i++ {
		if math.Abs(a[i]-m) < razn {
			razn = math.Abs(a[i] - m)
			elem = int(a[i])
		} else if math.Abs(a[i]-m) == razn && (int(a[i]) > elem) {
			elem = int(a[i])
		}
	}
	fmt.Println(elem)
}
