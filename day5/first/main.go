package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

const (
	DiagramSize int = 1000
)

func main() {
	in := input.Read("input.txt")

	r := regexp.MustCompile(`([\d]+),([\d]+) -> ([\d]+),([\d]+)`)

	diag := [DiagramSize][DiagramSize]int{}

	overlaps := 0

	for _, el := range in {
		m := r.FindStringSubmatch(el)

		x1, _ := strconv.Atoi(m[1])
		y1, _ := strconv.Atoi(m[2])
		x2, _ := strconv.Atoi(m[3])
		y2, _ := strconv.Atoi(m[4])

		if x1 == x2 {
			// horizontal line
			if y1 > y2 {
				// swap the variables so that we have a consistent for-loop
				y1, y2 = y2, y1
			}

			for i := y1; i <= y2; i++ {
				diag[i][x1]++

				if diag[i][x1] == 2 {
					// we've got overlapping lines, but only count first overlap
					overlaps++
				}
			}
		} else if y1 == y2 {
			// vertical line
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			for i := x1; i <= x2; i++ {
				diag[y1][i]++

				if diag[y1][i] == 2 {
					// we've got overlapping lines, but only count first overlap
					overlaps++
				}
			}
		}
	}

	fmt.Printf("Number of overlapping points %d\n", overlaps)
}
