package main

import (
	"fmt"
)

func main() {
	var buttonCount int
	fmt.Scan(&buttonCount)
	clickCount := make([]int, buttonCount)
	for i := 0; i < len(clickCount); i++ {
		fmt.Scan(&clickCount[i])
	}
	counterSort(clickCount)
}
func counterSort(list []int) {
	span := 100001
	counts := make([]int, span)
	var clickQuantity int
	fmt.Scan(&clickQuantity)
	dynSpan := 0
	for i := 0; i < clickQuantity; i++ {
		var sequenceClick int
		fmt.Scan(&sequenceClick)
		counts[sequenceClick]++
		if sequenceClick > dynSpan {
			dynSpan = sequenceClick
		}
	}
	span = dynSpan
	for i := 0; i < span; i++ {
		if counts[i+1] <= list[i] {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}
