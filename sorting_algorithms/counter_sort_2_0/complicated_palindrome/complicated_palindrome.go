package main

import (
	"fmt"
)

func main() {
	counterSort()
}
func counterSort() {
	var length int
	fmt.Scan(&length)
	var word1 string
	fmt.Scan(&word1)
	letters := make([]int, 26)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word1); j++ {
			if string(alphabet[i]) == string(word1[j]) {
				letters[i]++
			}
		}
	}
	oddLetter := ""
	for i := 0; i < len(letters); i++ {
		if letters[i]%2 != 0 || letters[i] == 1 {
			letters[i]--
			if oddLetter == "" {
				oddLetter = string(alphabet[i])
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < letters[i]/2; j++ {
			if letters[i] != 0 {
				fmt.Print(string(alphabet[i]))
			}
		}
	}
	fmt.Print(oddLetter)
	for i := 25; i >= 0; i-- {
		for j := 0; j < letters[i]/2; j++ {
			if letters[i] != 0 {
				fmt.Print(string(alphabet[i]))
			}
		}
	}
}
