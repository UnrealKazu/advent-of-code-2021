package paper

import (
	"reflect"
	"testing"
)

func TestFoldX_WithNoCollisionGrid_ShouldFoldCorrectly(t *testing.T) {
	// paper layout:
	/*
	 *	...
	 *	..X
	 *	...
	 *	..X
	 *
	 */
	p := Paper{
		Grid: map[int]map[int]bool{
			2: {
				1: true,
				3: true,
			},
		},
	}

	// fold over x = 1
	p.FoldX(1)

	// result should be:
	/*
	 *	.
	 *	X
	 *	.
	 *	X
	 *
	 */

	exp := map[int]map[int]bool{
		0: {
			1: true,
			3: true,
		},
	}

	if !reflect.DeepEqual(p.Grid, exp) {
		t.Errorf("Grid post fold is incorrect. Expected %v, got %v", exp, p.Grid)
	}
}

func TestFoldX_WithCollisionGrid_ShouldFoldCorrectly(t *testing.T) {
	// paper layout:
	/*
	 *	...
	 *	X.X
	 *	X..
	 *	..X
	 *
	 */
	p := Paper{
		Grid: map[int]map[int]bool{
			0: {
				1: true,
				2: true,
			},
			2: {
				1: true,
				3: true,
			},
		},
	}

	// fold over x = 1
	p.FoldX(1)

	// result should be:
	/*
	 *	.
	 *	X
	 *	X
	 *	X
	 *
	 */

	exp := map[int]map[int]bool{
		0: {
			1: true,
			2: true,
			3: true,
		},
	}

	if !reflect.DeepEqual(p.Grid, exp) {
		t.Errorf("Grid post fold is incorrect. Expected %v, got %v", exp, p.Grid)
	}
}

func TestFoldY_WithNoCollisionGrid_ShouldFoldCorrectly(t *testing.T) {
	// paper layout:
	/*
	 *	..X
	 *	...
	 *	...
	 *	..X
	 *  ...
	 */
	p := Paper{
		Grid: map[int]map[int]bool{
			2: {
				0: true,
				3: true,
			},
		},
	}

	// fold over y = 2
	p.FoldY(2)

	// result should be:
	/*
	 *	..X
	 *	..X
	 *
	 */

	exp := map[int]map[int]bool{
		2: {
			0: true,
			1: true,
		},
	}

	if !reflect.DeepEqual(p.Grid, exp) {
		t.Errorf("Grid post fold is incorrect. Expected %v, got %v", exp, p.Grid)
	}
}

func TestFoldY_WithNoCollisionGrid_ShouldRemoveFoldedPoints(t *testing.T) {
	// paper layout:
	/*
	 *	..X
	 *	...
	 *	...
	 *	..X
	 *  ...
	 */
	p := Paper{
		Grid: map[int]map[int]bool{
			2: {
				0: true,
				3: true,
			},
		},
	}

	// fold over y = 2
	p.FoldY(2)

	// result should be:
	/*
	 *	..X
	 *	..X
	 *
	 */

	if _, ex := p.Grid[2][3]; ex {
		t.Error("Folded point at coordinates x,y 2,3 still exists")
	}
}

func TestFoldY_WithCollisionGrid_ShouldFoldCorrectly(t *testing.T) {
	// paper layout:
	/*
	 *	..X
	 *	..X
	 *	...
	 *	..X
	 *  ...
	 */
	p := Paper{
		Grid: map[int]map[int]bool{
			2: {
				0: true,
				1: true,
				3: true,
			},
		},
	}

	// fold over y = 2
	p.FoldY(2)

	// result should be:
	/*
	 *	..X
	 *	..X
	 *
	 */

	exp := map[int]map[int]bool{
		2: {
			0: true,
			1: true,
		},
	}

	if !reflect.DeepEqual(p.Grid, exp) {
		t.Errorf("Grid post fold is incorrect. Expected %v, got %v", exp, p.Grid)
	}
}
