package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	result := 0

	for i := 0; i < 32; i++ {
		if (num&(1<<i))>>i == 1 {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
