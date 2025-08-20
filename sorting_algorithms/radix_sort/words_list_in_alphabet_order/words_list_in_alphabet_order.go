package main

import "fmt"

const (
	minVal     = 97  // a
	maxVal     = 122 // z
	pocketSize = maxVal - minVal + 2
)

func getMaxLen(slice []string) int {
	maxLen := 0
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) > maxLen {
			maxLen = len(slice[i])
		}
	}
	return maxLen
}

func radixSort(slice []string) {
	pocket := make([][]string, pocketSize)
	for i := 0; i < pocketSize; i++ {
		pocket[i] = make([]string, 0)
	}
	maxLen := getMaxLen(slice)
	for i := maxLen - 1; i >= 0; i-- {
		for j := 0; j < len(slice); j++ {
			var pocketIdx int
			if len(slice[j]) > i {
				pocketIdx = int(slice[j][i]) - minVal + 1
			}
			pocket[pocketIdx] = append(pocket[pocketIdx], slice[j])
		}
		sliceIdx := 0
		for j := 0; j < pocketSize; j++ {
			for k := 0; k < len(pocket[j]); k++ {
				slice[sliceIdx] = pocket[j][k]
				sliceIdx++
			}
			pocket[j] = pocket[j][:0]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	radixSort(slice)
	for i := 0; i < n; i++ {
		fmt.Println(slice[i])
	}
}
