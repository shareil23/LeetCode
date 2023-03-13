package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &n)

		var s string
		scanner.Scan()
		s = scanner.Text()

		fmt.Println(getParity(n, s))
	}
}

func getParity(n int, s string) string {
	res := ""
	for i := 1; i <= n; i++ {
		pi := s[:i]
		si := s[n-i:]

		ri := xor(pi, si)

		if countSetBits(ri)%2 == 0 {
			res += "0 "
		} else {
			res += "1 "
		}
	}
	return res
}

func xor(a, b string) string {
	res := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res += "0 "
		} else {
			res += "1 "
		}
	}
	return res
}

func countSetBits(s string) int {
	count := 0
	for _, ch := range s {
		if ch == '1' {
			count++
		}
	}
	return count
}
