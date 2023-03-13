package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	// Put into each matrix row additional values: [stength, index, 1 ... 1, 0 ... 0]
	lenC := len(mat[0])
	for r, row := range mat {
		strength := 0
		for strength < lenC && row[strength] == 1 {
			strength++
		}
		row[0] = strength
		row[1] = r
	}

	// Sort by: [ strength, index ]
	sort.Slice(mat, func(i, j int) bool {
		if mat[i][0] != mat[j][0] {
			return mat[i][0] <= mat[j][0]
		}
		return mat[i][1] <= mat[j][1]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = mat[i][1]
	}
	return result
}

func main() {
	var test_case = [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	fmt.Println(kWeakestRows(test_case, 3))
}
