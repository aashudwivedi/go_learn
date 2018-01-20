package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	var prev rune = ' '
	var bestSoFar, countSoFar int = 0, 0

	for i := range s {
		if rune(s[i]) == prev {
			countSoFar = 1
		} else {
			countSoFar += 1
		}
		if countSoFar > bestSoFar {
			bestSoFar = countSoFar
		}
		prev = rune(s[i])
	}
	return bestSoFar
}

func main() {
	var input = "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(input))
	fmt.Println(len(input))
}
