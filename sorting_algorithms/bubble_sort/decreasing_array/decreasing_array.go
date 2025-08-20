package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	a = bubbleSort(a)
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Print(a[i], " ")
	}
}
func bubbleSort(list []int) []int {
	for i := len(list) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
				flag = true
			}
		}
		if flag == false {
			return list
		}
	}
	return list
}
