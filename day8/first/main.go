package main

import (
	"fmt"
	"strings"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	count := 0

	for _, l := range in {
		r := strings.Split(l, " | ")[1]

		ds := strings.Split(r, " ")

		for _, d := range ds {
			// digits have the following length:
			// 1 -> 2
			// 4 -> 4
			// 7 -> 3
			// 8 -> 7
			switch len(d) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	fmt.Printf("Number of 1, 4, 7 and 8 digits %d\n", count)
}
