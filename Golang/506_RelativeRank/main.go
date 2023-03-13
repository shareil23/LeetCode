package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var reverse_score = make([]int, len(score))
  var temp_map = make(map[int]string)
  var temp_count int = 4
  var result []string

	copy(reverse_score, score)

	sort.Sort(sort.Reverse(sort.IntSlice(reverse_score)))

  for idx, i := range reverse_score {
    switch idx{
      case 0:
        temp_map[i] = "Gold Medal"
      
      case 1:
        temp_map[i] = "Silver Medal"

      case 2:
        temp_map[i] = "Bronze Medal"

      default:
        temp_map[i] = strconv.Itoa(temp_count)
        temp_count++
    }
  }

  for _, i := range score {
    result = append(result, temp_map[i])
  }

	return result
}

func main() {
	fmt.Println(findRelativeRanks([]int{1, 2, 3, 4, 5}))
}
