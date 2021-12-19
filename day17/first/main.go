package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

type Target struct {
	xMin, xMax int
	yMin, yMax int
}

type Vector struct {
	x, y int
}

func main() {
	in := input.Read("input.txt")

	r := regexp.MustCompile(`target area: x=([\-0-9]+)..([\-0-9]+), y=([\-0-9]+)..([\-0-9]+)`)

	rin := r.FindStringSubmatch(in[0])

	xMin, _ := strconv.Atoi(rin[1])
	xMax, _ := strconv.Atoi(rin[2])
	yMin, _ := strconv.Atoi(rin[4]) // careful to switch these around
	yMax, _ := strconv.Atoi(rin[3])

	t := NewTarget(xMin, xMax, yMin, yMax)

	y := findHighestY(t)

	fmt.Printf("Highest y vector is %d\n", y)
}

func findHighestY(t *Target) int {
	// x is always higher than 0, because we always go right
	// y is always higher than 0, because we want to find the highest y

	cands := []Vector{}

	for i := 0; i < t.xMax; i++ {
		if i > t.xMax {
			// overshot, no longer a feasible solution
			break
		}

		for j := 0; j < t.yMax*-1; j++ {
			if j < t.yMax {
				// we overshot, this is no longer a feasible solution
				break
			}

			if t.hitsTarget(i, j) {
				cands = append(cands, Vector{x: i, y: j})
			}
		}
	}

	// loop over all candidates
	highestY := 0

	for _, c := range cands {
		if c.y > highestY {
			highestY = c.y
		}
	}

	return (highestY * (highestY + 1)) / 2
}

func NewTarget(xMin, xMax, yMin, yMax int) *Target {
	return &Target{
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}
}

func (t *Target) hitsTarget(i, j int) bool {
	dx, dy := i, j

	x, y := dx, dy

	for x <= t.xMax && y >= t.yMax {
		if x >= t.xMin && x <= t.xMax && y <= t.yMin && y >= t.yMax {
			return true
		}

		if dx > 0 {
			dx--
		} else if dx < 0 {
			dx++
		}

		dy--

		x += dx
		y += dy
	}

	return false
}
