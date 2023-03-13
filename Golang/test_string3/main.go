package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(n int) [][]int {
	var result [][]int
	var temp_data []int

	for i := 1; i <= n; i++ {
		temp_data_2 := temp_data
		temp_data = nil
		counter_first, counter_second := 0, 1

		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				temp_data = append(temp_data, 1)
			} else {
				if len(temp_data_2) > 2 || counter_second <= len(temp_data_2)-1 {
					temp_data = append(temp_data, temp_data_2[counter_first]+temp_data_2[counter_second])
					counter_first++
					counter_second++
				}
			}
		}

		result = append(result, temp_data)
	}

	return result
}

func main() {
	result := solution(6)
	counter := 0
	fmt.Println(result)

	for i := len(result) - 1; i > 0; i-- {
		fmt.Print(strings.Repeat(" ", i))
		for _, y := range result[counter] {
			fmt.Print(strconv.Itoa(y) + " ")
		}

		fmt.Println()
		counter++
	}
}
