package main

import (
	"fmt"
)

func getLetters(word string) []int {
	letters := make([]int, 26)
	alphabet := "abcdefghijklmnopqrstuvwzyz"
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if string(alphabet[i]) == string(word[j]) {
				letters[i]++
			}
		}
	}
	return letters
}
func main() {
	var word1, word2 string
	fmt.Scan(&word1, &word2)
	letWords1 := getLetters(word1)
	letWords2 := getLetters(word2)
	for i := 0; i < 26; i++ {
		if letWords1[i] != letWords2[i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
