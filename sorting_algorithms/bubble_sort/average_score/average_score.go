package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	var nGds [][5]string
	for i := 0; i < number; i++ {
		var row [5]string
		fmt.Scan(&row[0], &row[1], &row[2], &row[3], &row[4])
		nGds = append(nGds, row)
	}
	nGds = bubbleSort(nGds)
	for _, row := range nGds {
		fmt.Println(row[0], row[1])
	}
}

func bubbleSort(list [][5]string) [][5]string {
	for i := len(list) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if helpAtoi(list[j][2], list[j][3], list[j][4]) < helpAtoi(list[j+1][2], list[j+1][3], list[j+1][4]) {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
		if flag == false {
			return list
		}
	}
	return list
}
func helpAtoi(k, j, c string) float64 {
	a, _ := strconv.Atoi(k)
	b, _ := strconv.Atoi(j)
	d, _ := strconv.Atoi(c)
	return float64(float64(a+b+d) / float64(3))
}
