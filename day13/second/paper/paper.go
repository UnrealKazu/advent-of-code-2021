// Package paper provides a paper struct on which folding operations can be performed
package paper

type Paper struct {
	Grid map[int]map[int]bool
}

func New() *Paper {
	return &Paper{
		Grid: make(map[int]map[int]bool),
	}
}

// AddDot adds a dot on the paper grid on the given x,y coordinates
func (p *Paper) AddDot(x, y int) {
	if _, ok := p.Grid[x]; !ok {
		p.Grid[x] = make(map[int]bool)
	}

	p.Grid[x][y] = true
}

// FoldX folds the paper sideways (to the left) over a given vertical line
// Any colliding points remain, the rest is merged and everything
// right of the fold is deleted
func (p *Paper) FoldX(fx int) {
	// fold every map on >X onto <X (mirrored), so fold to the left
	for x, m := range p.Grid {
		if x < fx {
			// leftside of the fold, not interested, moving on
			continue
		}

		// rightside of the fold, check if there is a collision map on the leftside of the fold
		cx := fx - (x - fx)

		cxMap, ex := p.Grid[cx]

		if ex {
			// there is a collision, we have to 'merge' the maps
			// (because dots remain, we can just apply all dots from the righthand side again)
			for y := range m {
				cxMap[y] = true
			}
		} else {
			// no collision map, so the entire righthand map is placed here
			p.Grid[cx] = m
		}

		// now we can delete the original map from the righthand side
		delete(p.Grid, x)
	}
}

// FoldY folds the paper upwards over a given horizontal line
// Any colliding points remain, the rest is merged and everything
// below the fold is deleted
func (p *Paper) FoldY(fy int) {
	// loop over all x-maps first
	for _, mx := range p.Grid {
		// next, check all y maps to see if they are foldable
		for y := range mx {
			if y < fy {
				// above the fold, skip
				continue
			}

			// calculate the new y position above the fold
			ny := fy - (y - fy)

			// set the new position
			mx[ny] = true

			// delete the dot below the fold
			delete(mx, y)
		}
	}
}

// ToString converts the paper grid into a human readable string
func (p *Paper) ToString() string {
	// first, get to know the biggest X and Y
	maxY := 0
	maxX := 0
	for x, mx := range p.Grid {
		if x > maxX {
			maxX = x
		}

		for y := range mx {
			if y > maxY {
				maxY = y
			}
		}
	}

	m := [][]string{}

	// initialize the 2-dimensional slice, makes it easier later on to reference directly
	for y := 0; y <= maxY; y++ {
		m = append(m, []string{})

		for x := 0; x <= maxX; x++ {
			m[y] = make([]string, maxX+1)
		}
	}

	// important to keep in mind that the grid is inverted with respect to the 2-dimensional slice
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, ex1 := p.Grid[x]; ex1 {
				// the x map exists, check the y
				if _, ex2 := p.Grid[x][y]; ex2 {
					// y also exists, so a dot (a dot is actually a #, a lack of dot is .)
					// confusing, I know, but I didn't design the puzzles...
					m[y][x] = "#"
				} else {
					// does not exist, so a dot
					m[y][x] = "."
				}
			} else {
				// does not exist, so no dot
				m[y][x] = "."
			}
		}
	}

	// finally, loop over the 2D slice and convert it into a single string with newlines
	ret := ""
	for _, row := range m {
		for _, col := range row {
			ret += col
		}
		ret += "\n"
	}

	return ret
}
