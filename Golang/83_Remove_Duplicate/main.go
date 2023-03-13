package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var tmp map[int]int = make(map[int]int)
	var tmp_unique []int
	var result *ListNode

	for head != nil {
		if _, is_exist := tmp[head.Val]; is_exist {
			head = head.Next
			continue
		}

		tmp[head.Val]++
		tmp_unique = append(tmp_unique, head.Val)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tmp_unique)))

	for _, i := range tmp_unique {
		result = &ListNode{i, result}
	}

	return result
}

func deleteDuplicates2(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

func main() {
	var test_var *ListNode

	for _, i := range []int{3, 3, 2, 1, 1} {
		test_var = &ListNode{
			Val:  i,
			Next: test_var,
		}
	}

	fmt.Println(deleteDuplicates(test_var))
	fmt.Println(deleteDuplicates2(test_var))
}
