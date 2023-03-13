package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var index int
	var target_index int = 1
	var result int
	var temp_string string
	var status bool = false

	if s == " " || len(s) == 1 {
		return 1
	} else if s == "" {
		return 0
	}

	for _, i := range s {
		if strings.Contains(temp_string, string(i)) {
			if len(temp_string) > result {
				result = len(temp_string)
			}

			index = target_index - 1
			status = true
		}

		temp_string = s[index:target_index]
		fmt.Println(temp_string)

		if target_index > len(s) {
			break
		}

		target_index += 1
	}

	if !status {
		result = len(temp_string)
	}

	return result
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)
	res := 0

	if s == "" {
		return 0
	}

	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			l = max(l, index)
		}
		res = max(res, r-l+1)
		m[s[r]] = r + 1
	}
	return res
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func main() {
	var test_string string = "aab"
	fmt.Println(lengthOfLongestSubstring(test_string))
}
