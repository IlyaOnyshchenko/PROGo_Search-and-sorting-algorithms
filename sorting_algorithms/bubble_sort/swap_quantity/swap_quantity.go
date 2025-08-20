package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	slice := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Print(bubbleSort(slice))
}
func bubbleSort(list []int) int {
	cnt := 0
	for i := len(list) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
				flag = true
				cnt++
			}
		}
		if flag == false {
			return cnt
		}
	}
	return cnt
}
