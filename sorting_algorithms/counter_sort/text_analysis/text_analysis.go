package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var number int
	fmt.Scan(&number)
	reader := bufio.NewReader(os.Stdin)
	words := make([]string, number)
	for i := 0; i <= number; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		wordsInLine := strings.Fields(input)
		words = append(words, wordsInLine...)
	}
	words = removeEmptyStrings(words)
	counterSort(words, len(words))
}
func removeEmptyStrings(s []string) []string {
	var result []string
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func counterSort(list []string, width int) {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	span := 26
	dynSpan := 0
	counts := make([]int, 26)
	for i := 0; i < width; i++ {
		counts[strings.Index(alphabet, string(list[i][0]))]++
		if strings.Index(alphabet, string(list[i][0])) > dynSpan {
			dynSpan = strings.Index(alphabet, string(list[i][0]))
		}
	}
	span = dynSpan
	for i := 0; i < span+1; i++ {
		if counts[i] != 0 {
			fmt.Println(string(alphabet[i]), counts[i])
		}
	}
}
