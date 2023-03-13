package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	var temp = make(map[string]int)

	for _, r := range s {
		temp[string(r)]++
	}

	var result int

	for _, i := range temp {
		result += i / 2 * 2

		if result%2 == 0 && i%2 == 1 {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
