package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/advent-of-code-2021/day15/second/graph"
	"github.com/advent-of-code-2021/input"
)

const (
	MultiplyRateX = 5
	MultiplyRateY = 5
)

func main() {
	start := time.Now()
	in := input.Read("input.txt")

	nin := parseInput(in)

	mnin := multiplyMap(nin, 5, 5)

	p := GetLowestRiskPath(mnin)

	fmt.Printf("The lowest total risk path is %d\n", p)
	fmt.Printf("Calculation time: %s\n", time.Since(start))
}

func GetLowestRiskPath(in [][]int) int {
	t := graph.New(in)

	return t.GetLowestRiskPath()
}

func parseInput(in []string) [][]int {
	ret := [][]int{}

	for i, row := range in {
		ret = append(ret, []int{})

		for _, col := range row {
			v, _ := strconv.Atoi(string(col))
			ret[i] = append(ret[i], v)
		}
	}

	return ret
}

func multiplyMap(m [][]int, mpX, mpY int) [][]int {
	// first, duplicate (and increase) everything horizontally
	// iterate over all rows
	for i := 0; i < len(m); i++ {
		s := len(m[i])
		// iterate over the number of tiles we need to create
		for a := 1; a < mpX; a++ {
			// and finally, iterate over all columns
			for j := 0; j < s; j++ {
				// get the corresponding value from the previous tile
				prevJ := m[i][((a-1)*s)+j]

				// calculate the new value
				newJ := (prevJ % 9) + 1

				m[i] = append(m[i], newJ)
			}
		}
	}

	// next, duplicate (and increase) everything vertically
	s := len(m)
	for b := 1; b < mpY; b++ {
		for i := 0; i < s; i++ {
			// add a new row to the slice
			nrow := []int{}

			for j := 0; j < len(m[i]); j++ {
				// get the corresponding value from the previous tile
				prev := m[((b-1)*s)+i][j]

				// calculate the new value
				new := (prev % 9) + 1

				nrow = append(nrow, new)
			}

			m = append(m, nrow)
		}
	}

	return m
}
