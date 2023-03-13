package main

import "fmt"

func numberOfSteps(num int) int {
	var result int

	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		result += 1
	}

	return result
}

func main() {
	fmt.Println(numberOfSteps(8))
}
