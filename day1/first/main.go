package main

import (
	"fmt"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	count := 0
	for i := 1; i < len(in); i++ {
		prev, _ := strconv.Atoi(in[i-1])
		cur, _ := strconv.Atoi(in[i])

		if cur-prev > 0 {
			count++
		}
	}

	fmt.Printf("Number of increases: %d\n", count)
}
