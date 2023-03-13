package main

import "fmt"

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func solution(list1 *LinkedNode, list2 *LinkedNode, a int, b int) *LinkedNode {
	var result *LinkedNode
	var tmp_arr_1, tmp_arr_2 []int
	var list1_counter int = 1
	var status bool = true

	for list2 != nil {
		tmp_arr_2 = append(tmp_arr_2, list2.Val)
		list2 = list2.Next
	}

	for list1 != nil {
		if list1_counter >= a && list1_counter <= b {
			list1 = list1.Next
			list1_counter++
			continue
		}

		tmp_arr_1 = append(tmp_arr_1, list1.Val)
		if status {
			tmp_arr_1 = append(tmp_arr_1, tmp_arr_2...)
			status = false
		}

		list1 = list1.Next
		list1_counter++
	}

	fmt.Println(tmp_arr_1)

	for i := len(tmp_arr_1) - 1; i >= 0; i-- {
		result = &LinkedNode{
			Val:  i,
			Next: result,
		}
	}

	return result
}

func main() {
	var list1 *LinkedNode
	var list2 *LinkedNode
	var a int = 3
	var b int = 4

	for _, i := range []int{4, 3, 2, 1} {
		list1 = &LinkedNode{
			Val:  i,
			Next: list1,
		}
	}

	for _, i := range []int{14, 13, 12, 11} {
		list2 = &LinkedNode{
			Val:  i,
			Next: list2,
		}
	}

	fmt.Println(solution(list1, list2, a, b))
}
