package main

import "fmt"

func main() {
	k := 1001
	count := make([]int, k)
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		count[x] = count[x] + 1
	}

	var res int
	for i := 1; i < k; i++ {
		if count[i] == 2 {
			res++
		}
	}

	fmt.Println(res)
}
