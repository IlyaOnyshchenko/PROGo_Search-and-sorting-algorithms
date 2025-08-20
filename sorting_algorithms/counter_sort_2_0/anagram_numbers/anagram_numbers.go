package main

import (
	"fmt"
)

func main() {
	fmt.Println(counterSort())
}
func counterSort() string {
	var number1, number2 int
	fmt.Scan(&number1, &number2)
	span := 1000000001
	counts1 := make([]int, span)
	counts2 := make([]int, span)
	span = 0
	span1 := 0
	for i := 0; i < len(divNum(number1)); i++ {
		counts1[divNum(number1)[i]]++
		if divNum(number1)[i] > span {
			span = divNum(number1)[i]
		}
	}
	for i := 0; i < len(divNum(number2)); i++ {
		counts2[divNum(number2)[i]]++
		if divNum(number2)[i] > span1 {
			span1 = divNum(number2)[i]
		}
	}
	counts1 = counts1[:span+1]
	counts2 = counts2[:span1+1]
	for i := 0; i < len(counts1); i++ {
		if counts1[i] != counts2[i] {
			return "NO"
		}
	}
	return "YES"
}

func divNum(num int) []int {
	var numbers []int
	for num > 0 {
		numbers = append(numbers, num%10)
		num /= 10
	}
	return numbers
}
