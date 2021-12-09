package main

import (
	"fmt"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	hm := parseInput(in)

	riskSum := getRiskSum(hm)

	fmt.Printf("Sum of all risk levels is %d\n", riskSum)
}

func getRiskSum(hm [][]int) int {
	sum := 0
	for i := 0; i < len(hm); i++ {
		for j := 0; j < len(hm[i]); j++ {
			if checkIfLowest(i, j, hm) {
				sum += 1 + hm[i][j]
			}
		}
	}

	return sum
}

// checkIfLowest checks the neighbours of a point in all four directions (if present),
// and returns true if the point at the given coordinate is the lowest of all four neighbours
func checkIfLowest(i, j int, hm [][]int) bool {
	cur := hm[i][j]

	// left neigbour
	if j > 0 && cur >= hm[i][j-1] {
		return false
	}

	// top neighbour
	if i > 0 && cur >= hm[i-1][j] {
		return false
	}

	// right neigbour
	if j < len(hm[i])-1 && cur >= hm[i][j+1] {
		return false
	}

	// bottom neighbour
	if i < len(hm)-1 && cur >= hm[i+1][j] {
		return false
	}

	return true
}

// parseInput simply converts the slice of strings into a two-dimensional int slice
func parseInput(in []string) [][]int {
	hm := [][]int{}

	for _, l := range in {
		row := []int{}
		for _, r := range l {
			conv, _ := strconv.Atoi(string(r))

			row = append(row, conv)
		}

		hm = append(hm, row)
	}

	return hm
}
