package main

import "fmt"

var (
	fibonacciCache = make(map[int]int)
)

func climbStairs(n int) int {
	/**
	this solution is using caching method
	to reduce the high number of fibonacci recursive
	*/

	if n <= 2 {
		return n
	}

	if result, ok := fibonacciCache[n]; ok {
		return result
	}

	result := climbStairs(n-1) + climbStairs(n-2)
	fibonacciCache[n] = result

	return result
}

func climbStairs2(n int) int {
	/**
	this solution is using iterative method
	*/

	dp := make([]int, n+2)

	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

func main() {
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs2(5))
}
