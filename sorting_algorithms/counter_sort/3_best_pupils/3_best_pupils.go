package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	var namGrs [][3]string
	for i := 0; i < number; i++ {
		var row [3]string
		fmt.Scan(&row[0], &row[1], &row[2])
		namGrs = append(namGrs, row)
	}
	counterSort(namGrs, number)
}
func counterSort(list [][3]string, width int) {
	var span int = 1000001
	cnt := 0
	dynSpan := 0
	counts := make([]int, span)
	for i := 0; i < width; i++ {
		counts[helpAtoi(list[i][2])]++
		if helpAtoi(list[i][2]) > dynSpan {
			dynSpan = helpAtoi(list[i][2])
		}
	}
	span = dynSpan
	for i := span; i >= 0; i-- {
		for j := 0; j < width; j++ {
			if helpAtoi(list[j][2]) == i {
				fmt.Println(list[j][0] + " " + list[j][1])
				cnt++
			}
		}
		if cnt == 3 {
			break
		}
	}
}
func helpAtoi(k string) int {
	a, _ := strconv.Atoi(k)
	return a
}
