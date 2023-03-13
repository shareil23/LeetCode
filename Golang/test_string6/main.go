package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var score int
	for i := 0; i < n/2; i++ {
		num1, num2 := nums[i*2], nums[i*2+1]
		bit := num1 & num2 & -num1 & -num2
		score += bit
	}

	fmt.Println(score)
}
