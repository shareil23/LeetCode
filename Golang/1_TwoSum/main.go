package main

import "fmt"

func solution(nums []int, target int) []int {
	// this is native / bruteforce approach

	var result []int
	var temp_1, temp_2 int

	for idx_1, i := range nums {
		for idx_2, j := range nums {
			if i+j == target && idx_1 != idx_2 {
				temp_1 = idx_1
				temp_2 = idx_2
			}
		}
	}

	result = append(result, temp_1)
	result = append(result, temp_2)

	return result
}

func solution2(nums []int, target int) []int {
	/**
	The map key will be fill by value of current index
	and the value will be fill with current index
	so we will check if the target - curr value are exists in map
	if exists will take the idx that exists in map
	*/

	temp_map := make(map[int]int)

	for index, value := range nums {
		if idx, is_ok := temp_map[target-value]; is_ok {
			return []int{idx, index}
		}

		temp_map[value] = index
	}

	return []int{}
}

func main() {
	fmt.Println(solution2(
		[]int{2, 7, 11, 15},
		9,
	))
}
