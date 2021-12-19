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

func main() {
	in := input.Read("input.txt")

	r := regexp.MustCompile(`target area: x=([\-0-9]+)..([\-0-9]+), y=([\-0-9]+)..([\-0-9]+)`)

	rin := r.FindStringSubmatch(in[0])

	xMin, _ := strconv.Atoi(rin[1])
	xMax, _ := strconv.Atoi(rin[2])
	yMin, _ := strconv.Atoi(rin[4]) // careful to switch these around
	yMax, _ := strconv.Atoi(rin[3])

	t := NewTarget(xMin, xMax, yMin, yMax)

	count := findHitCount(t)

	fmt.Printf("Highest number of hit velocity vectors is %d\n", count)
}

func findHitCount(t *Target) int {
	count := 0

	for i := 0; i <= t.xMax; i++ {
		if i > t.xMax {
			// overshot, no longer a feasible solution
			break
		}

		for j := t.yMax; j <= t.yMax*-1; j++ {
			if j < t.yMax {
				// we overshot, this is no longer a feasible solution
				break
			}

			if t.hitsTarget(i, j) {
				count++
			}
		}
	}

	return count
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
