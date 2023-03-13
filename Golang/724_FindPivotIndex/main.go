package main

import "fmt"

func slice_sum(nums []int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}

func pivotIndex(nums []int) int {
	var result int = -1
	var left_val, right_val = 0, 0

	for idx, _ := range nums {
		if idx == 0 {
			right_val = slice_sum(nums[idx+1:])

			if left_val == right_val {
				result = idx
				break
			}

			continue
		}

		if idx == len(nums)-1 {
			right_val = 0
			left_val = slice_sum(nums[:len(nums)-1])

			if left_val == right_val {
				result = idx
				break
			}

			break
		}

		right_val = slice_sum(nums[:idx])
		left_val = slice_sum(nums[idx+1:])

		if left_val == right_val {
			result = idx
			break
		}

		right_val = 0
		left_val = 0
	}

	return result
}

func pivotIndex2(nums []int) int {
	rightSum := slice_sum(nums)
	leftSum := 0

	for i, v := range nums {
		rightSum = rightSum - v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}

	return -1
}

func main() {
	var result []int = []int{-1, -1, 0, 0, -1, -1}

	fmt.Println(pivotIndex(result))
}
