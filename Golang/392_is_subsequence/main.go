package main

import (
	"fmt"
	"strings"
)

func findIndex(s []string, target string) (bool, int) {
	for idx, val := range s {
		if val == target {
			return true, idx
		}
	}

	return false, 0
}

func isSubsequence(s string, t string) bool {
	var t_split []string = strings.Split(t, "")
	var index_temp int

	for _, i := range s {
		idx_status, get_indexes := findIndex(t_split, string(i))

		fmt.Println("i: ", string(i), " get_indexes: ", get_indexes, " index_temp: ", index_temp)
		fmt.Println("result: ", get_indexes >= index_temp)

		if idx_status && get_indexes >= index_temp {
			index_temp = get_indexes
			t_split[get_indexes] = "_"
		} else {
			return false
		}
	}

	return true
}

func main() {
	//fmt.Println(isSubsequence("abc", "ahbgdc"))
	//fmt.Println(isSubsequence("axc", "ahbgdc"))
	//fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
	fmt.Println(isSubsequence("ab", "baab"))
}
