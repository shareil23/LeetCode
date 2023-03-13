package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	var result bool = true
	var temp_map = make(map[string]int)
	var temp_map2 = make(map[string]int)

	for _, i := range ransomNote {
		temp_map[string(i)] += 1
	}

	for _, i := range magazine {
		temp_map2[string(i)] += 1
	}

	for key, val := range temp_map {
		if temp_map2[key]-val < 0 {
			result = false
		}
	}

	return result
}

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}
