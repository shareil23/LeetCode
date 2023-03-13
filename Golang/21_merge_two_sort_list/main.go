package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var result *ListNode
	var temp_list []int

	for list1 != nil {
		temp_list = append(temp_list, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		temp_list = append(temp_list, list2.Val)
		list2 = list2.Next
	}

	sort.Sort(sort.Reverse(sort.IntSlice(temp_list)))

	for _, i := range temp_list {
		result = &ListNode{
			Val:  i,
			Next: result,
		}
	}

	return result
}
