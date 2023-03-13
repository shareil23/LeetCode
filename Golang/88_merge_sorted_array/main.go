package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n != 0 {
		copy(nums1[m:], nums2)
	}

	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println(m)

	if n != 0 {
		counter := 0

		for idx, i := range nums1 {
			if i == 0 && counter < len(nums2) {
				nums1[idx] = nums2[counter]
				counter++
			}
		}

		counter = 0
	}

	sort.Ints(nums1)
}

func main() {

}
