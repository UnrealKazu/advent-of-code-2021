package main

import (
	"fmt"
	"strconv"

	"github.com/advent-of-code-2021/day15/first/graph"
	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	nin := parseInput(in)

	p := GetLowestRiskPath(nin)

	fmt.Printf("The lowest total risk path is %d\n", p)
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
