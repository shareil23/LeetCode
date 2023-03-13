package main

import "fmt"

func getN(n_case int) {
	if n_case <= 0 {
		return
	}

	var n int

	// get the n number will be to sum
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	fmt.Println("\n", calculateSum(n))

	getN(n_case - 1)
}

func calculateSum(n int) int {
	if n == 0 {
		return 0
	}

	var number int

	// get the number will be to sum
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return 0
	}

	if number > 0 {
		return number*number + calculateSum(n-1)
	}

	return calculateSum(n - 1)
}

func main() {
	var n_case int

	// get how many test cases
	_, err := fmt.Scanf("%d", &n_case)
	if err != nil {
		return
	}

	getN(n_case)
}
