package main

import (
	"fmt"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.ReadInt("input.txt")

	count := 0
	for i := 1; i < len(in)-2; i++ {
		prev := in[i-1] + in[i] + in[i+1]
		cur := in[i] + in[i+1] + in[i+2]

		if cur-prev > 0 {
			count++
		}
	}

	fmt.Printf("Number of increases: %d\n", count)
}
