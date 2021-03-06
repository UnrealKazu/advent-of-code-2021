package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

type Point struct {
	i, j int
}

func main() {
	in := input.Read("input.txt")

	hm := parseInput(in)

	lPoints := getLowestPoints(hm)

	sizes := []int{}
	for _, p := range lPoints {
		sizes = append(sizes, getBasinSize(p, &hm))
	}

	sort.Ints(sizes)

	fmt.Printf("Sizes of the three biggest basins multiplied is %d\n", sizes[len(sizes)-1]*sizes[len(sizes)-2]*sizes[len(sizes)-3])
}

// getBasinSize returns the size of the basin (surprise) of the given low point. It traverses all neighbours and
// counts all non boundary ones
func getBasinSize(p Point, hm *[][]int) int {
	count := 0
	unchecked := []Point{}

	(*hm)[p.i][p.j] = 10
	count++

	// fill the unchecked with a first round
	unchecked = append(unchecked, getUncheckedNonBoundaryNeighbours(p.i, p.j, hm)...)
	count += len(unchecked)

	// continue adding more neighbours until we have nothing more to check
	for len(unchecked) > 0 {
		cur := unchecked[0]

		newPoints := getUncheckedNonBoundaryNeighbours(cur.i, cur.j, hm)
		count += len(newPoints)

		// remove the just checked point (the first), and add the newly discovered once
		unchecked = append(unchecked[1:], newPoints...)
	}

	// do a count on all neighbours
	return count
}

// getUncheckedNonBoundaryNeighbours returns a slice of points that are not a boundary point (i.e. not a 9),
// and have not been traversed before (i.e. a 10)
func getUncheckedNonBoundaryNeighbours(i, j int, hm *[][]int) []Point {
	unchecked := []Point{}
	// left neigbour
	if j > 0 && (*hm)[i][j-1] < 9 {
		unchecked = append(unchecked, Point{
			i: i,
			j: j - 1,
		})
		(*hm)[i][j-1] = 10
	}

	// top neighbour
	if i > 0 && (*hm)[i-1][j] < 9 {
		unchecked = append(unchecked, Point{
			i: i - 1,
			j: j,
		})
		(*hm)[i-1][j] = 10
	}

	// right neigbour
	if j < len((*hm)[i])-1 && (*hm)[i][j+1] < 9 {
		unchecked = append(unchecked, Point{
			i: i,
			j: j + 1,
		})
		(*hm)[i][j+1] = 10
	}

	// bottom neighbour
	if i < len((*hm))-1 && (*hm)[i+1][j] < 9 {
		unchecked = append(unchecked, Point{
			i: i + 1,
			j: j,
		})
		(*hm)[i+1][j] = 10
	}

	return unchecked
}

// getLowestPoints returns a slice of all points that are the lowest value of all their neighbours
func getLowestPoints(hm [][]int) []Point {
	points := []Point{}
	for i := 0; i < len(hm); i++ {
		for j := 0; j < len(hm[i]); j++ {
			if checkIfLowest(i, j, hm) {
				points = append(points, Point{
					i: i,
					j: j,
				})
			}
		}
	}

	return points
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
