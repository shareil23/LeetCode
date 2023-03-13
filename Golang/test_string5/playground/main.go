package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func squareSum(testcase []string) int {
	i, err := strconv.Atoi(testcase[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rest := testcase[1:]
	square := i * i

	switch {
	case i < 0 && len(rest) == 0:
		return 0
	case i < 0 && len(rest) > 0:
		return 0 + squareSum(rest)
	case i >= 0 && len(rest) == 0:
		return square
	case i >= 0 && len(rest) > 0:
		return square + squareSum(rest)
	}

	return 0
}

func printArray(sumArray []int) {
	rest := sumArray[1:]
	fmt.Println(sumArray[0])

	if len(rest) == 0 {
		return
	} else {
		printArray(rest)
	}

	return
}

func scanInput(checkNum int, numTestCases int, sumArray []int, Scanner *bufio.Scanner) {
	Scanner.Scan()
	input := Scanner.Text()

	if checkNum == 0 {
		firstInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		} else if err == nil && firstInput == 0 {
			os.Exit(0)
		}

		numTestCases = firstInput
	}

	result := strings.Split(input, " ")
	if checkNum > 0 && checkNum%2 == 0 {
		sum := squareSum(result)
		sumArray = append(sumArray, sum)
	}

	if numTestCases == checkNum/2 {
		printArray(sumArray)
		return
	} else {
		checkNum += 1
		scanInput(checkNum, numTestCases, sumArray, Scanner)
	}

	return
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	checkNum := 0
	sumArray := []int{}
	numTestCases := 0

	scanInput(checkNum, numTestCases, sumArray, Scanner)
}
