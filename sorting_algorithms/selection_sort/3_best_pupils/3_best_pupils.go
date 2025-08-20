package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	var nGds [][3]string
	for i := 0; i < number; i++ {
		var row [3]string
		fmt.Scan(&row[0], &row[1], &row[2])
		nGds = append(nGds, row)
	}
	nGds = selectionSort(nGds)
	for i := 0; i < 3; i++ {
		fmt.Println(nGds[i][0], nGds[i][1])
	}
}
func selectionSort(list [][3]string) [][3]string {
	for i := 0; i < 3; i++ {
		maxIndex := i
		for j := len(list) - 1; j > i; j-- {
			if helpAtoi(list[j][2]) > helpAtoi(list[maxIndex][2]) {
				maxIndex = j
			}
		}
		if maxIndex != i {
			temp := list[maxIndex]
			list[maxIndex] = list[i]
			list[i] = temp
		}
	}
	return list
}
func helpAtoi(k string) float64 {
	a, _ := strconv.Atoi(k)
	return float64(a)
}
