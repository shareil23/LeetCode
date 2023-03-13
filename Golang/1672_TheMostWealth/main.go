package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {
	var result int

	for _, i := range accounts {
		tempResult := 0

		for _, j := range i {
			tempResult += j
		}

		if tempResult > result {
			result = tempResult
		}
	}

	return result
}

func main() {
	//var testCase = [][]int{
	//	{1, 2, 3},
	//	{3, 2, 1},
	//}

	var testCase = [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	fmt.Println(maximumWealth(testCase))
}
