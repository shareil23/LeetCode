package main

import "fmt"

func Solution(AA int, AB int, BB int) string {
	// Implement your solution here
	var max_number int
	var selected_loop, result string

	if AA > max_number {
		max_number = AA
		selected_loop = "AA"
	}

	if AB > max_number {
		max_number = AB
		selected_loop = "AB"
	}

	if BB > max_number {
		max_number = BB
		selected_loop = "BB"
	}

	switch selected_loop {
	case "AA":
		
		for max_number != 0 {
			if BB != 0 {
				result += "AA"
			}
		}

	}
}

func main() {
	fmt.Println(Solution(5, 0, 2))
}
