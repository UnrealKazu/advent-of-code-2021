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

	fmt.Printf("Sizes of the three biggest basins multiplied is %d\n", sizes[len(sizes)-1]*sizes[len(sizes)-2]*sizes[len(sizes)-2])
}

func getBasinSize(p Point, hm *[][]int) int {
	neighs := make(map[int]map[int]int)
	unchecked := []Point{}

	addToMap(p.i, p.j, neighs)
	(*hm)[p.i][p.j] = 10

	// fill the unchecked with a first round
	getNonBoundaryNeighbours(p.i, p.j, hm, neighs, &unchecked)

	for len(unchecked) > 0 {
		cur := unchecked[0]

		getNonBoundaryNeighbours(cur.i, cur.j, hm, neighs, &unchecked)

		// remove the just checked point
		unchecked = unchecked[1:]
	}

	// do a count on all neighbours
	count := 0

	for _, i := range neighs {
		count += len(i)
	}

	return count
}

func getNonBoundaryNeighbours(i, j int, hm *[][]int, neighs map[int]map[int]int, unchecked *[]Point) {
	// left neigbour
	if j > 0 && (*hm)[i][j-1] < 9 {
		if addToMap(i, j-1, neighs) {
			*unchecked = append(*unchecked, Point{
				i: i,
				j: j - 1,
			})
		}
		(*hm)[i][j-1] = 10
	}

	// top neighbour
	if i > 0 && (*hm)[i-1][j] < 9 {
		if addToMap(i-1, j, neighs) {
			*unchecked = append(*unchecked, Point{
				i: i - 1,
				j: j,
			})
		}
		(*hm)[i-1][j] = 10
	}

	// right neigbour
	if j < len((*hm)[i])-1 && (*hm)[i][j+1] < 9 {
		if addToMap(i, j+1, neighs) {
			*unchecked = append(*unchecked, Point{
				i: i,
				j: j + 1,
			})
		}
		(*hm)[i][j+1] = 10
	}

	// bottom neighbour
	if i < len((*hm))-1 && (*hm)[i+1][j] < 9 {
		if addToMap(i+1, j, neighs) {
			*unchecked = append(*unchecked, Point{
				i: i + 1,
				j: j,
			})
		}
		(*hm)[i+1][j] = 10
	}
}

func addToMap(i, j int, neighs map[int]map[int]int) bool {
	if _, ok := neighs[i]; ok {
		// i is in the map
		if _, inOk := neighs[i][j]; inOk {
			// j is also in the map, nothing to add
			return false
		}
	} else {
		// i doesn't exist yet, so make it so we can add it
		neighs[i] = make(map[int]int)
	}

	neighs[i][j] = 1
	return true
}

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
