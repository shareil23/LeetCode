package main

import (
	"fmt"
	"math"
	"strconv"
)

func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	parseN := strconv.Itoa(n)
	status := true
	counter := 1

	for status {
		tmpSum := 0

		for _, i := range parseN {
			tmpSum += int(math.Pow(float64(i-48), 2))
		}

		if len(strconv.Itoa(tmpSum)) == 1 && counter > 1 {
			if strconv.Itoa(tmpSum) == "1" {
				return true
			} else {
				return false
			}
		}

		parseN = strconv.Itoa(tmpSum)
		counter++
	}

	return false
}

func main() {
	fmt.Println(isHappy(1111111))
}
