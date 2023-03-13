package main

import "fmt"

func singleNumber(nums []int) int {
	var tmp_map map[int]int = make(map[int]int)

	for _, i := range nums {
		tmp_map[i]++
	}

	for key, val := range tmp_map {
		if val == 1 {
			return key
		}
	}

	return 0
}

func singleNumber2(nums []int) int {
	var result int

	for _, i := range nums {
		result ^= i
	}

	return result
}

func singleNumber3(nums []int) int {
	var tmp_map map[int]int = make(map[int]int)
	var result int

	for _, i := range nums {
		if _, ok := tmp_map[i]; ok {
			result -= i
			continue
		}

		result += i
		tmp_map[i]++
	}

	return result
}

func main() {
	var test []int = []int{4, 1, 2, 1, 2}

	fmt.Println("result 1:", singleNumber(test))
	fmt.Println("result 2:", singleNumber2(test))
}
