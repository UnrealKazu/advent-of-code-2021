package main

import (
	"fmt"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

const (
	Debug = false
)

type Cavern struct {
	Step                int
	PreviousNrofFlashes int
	NrofFlashes         int
	Board               [][]int
	BoardSize           int
}

func main() {
	c := newCavern(input.Read("input.txt"))

	// keep simulating until we have a synchronized step
	for {
		c.IncreaseStep()

		if c.NrofFlashes-c.PreviousNrofFlashes == c.BoardSize {
			fmt.Printf("Synchronized flash detected at step %d\n", c.Step)
			break
		}

		c.PreviousNrofFlashes = c.NrofFlashes
	}
}

func (c *Cavern) Flash(i, j int) {
	if c.Board[i][j] <= 9 {
		panic("Unexpected flash, logic is flawed")
	}

	// first, set the flashed one to -1
	c.Board[i][j] = -1
	// and register the flash
	c.NrofFlashes++

	// next, go through all neighbours to see if they need to flash as well
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if k == 0 && l == 0 {
				// skip this combination, because this is the originally flashed octopus
				continue
			}

			if i+k < 0 || i+k >= len(c.Board) {
				// either above the top of the board, or below the bottom of the board
				// so this is an invalid neighbour, continue
				continue
			}

			if j+l < 0 || j+l >= len(c.Board[i+k]) {
				// either to the left or right of the board
				// so this is an invalid neighbour, continue
				continue
			}

			// increase the energy level of this neighbour, but only if it didn't flash before
			if c.Board[i+k][j+l] != -1 {
				c.Board[i+k][j+l]++

				if c.Board[i+k][j+l] > 9 {
					// this neighbour needs to flash as well
					c.Flash(i+k, j+l)
				}
			}
		}
	}
}

// IncreaseStep increases the board step and loops through all octopi and increase their energy level
func (c *Cavern) IncreaseStep() {
	c.Step++

	// first, increase all energy levels
	for i := 0; i < len(c.Board); i++ {
		for j := 0; j < len(c.Board[i]); j++ {
			if c.Board[i][j] == -1 {
				// special 'already flashed' case, which in essence is 0,
				// so an increase is 1
				c.Board[i][j] = 1
			} else {
				c.Board[i][j]++
			}
		}
	}

	// next, loop again to check for flashes
	for i := 0; i < len(c.Board); i++ {
		for j := 0; j < len(c.Board[i]); j++ {
			if c.Board[i][j] > 9 {
				c.Flash(i, j)
			}
		}
	}
}

// newCavern simply converts the slice of strings into a two-dimensional int slice
// and returns a new Cavern struct
func newCavern(in []string) Cavern {
	hm := [][]int{}

	for _, l := range in {
		row := []int{}
		for _, r := range l {
			conv, _ := strconv.Atoi(string(r))

			row = append(row, conv)
		}

		hm = append(hm, row)
	}

	return Cavern{
		Step:                0,
		PreviousNrofFlashes: -1,
		NrofFlashes:         0,
		Board:               hm,
		BoardSize:           len(hm) * len(hm[0]),
	}
}
