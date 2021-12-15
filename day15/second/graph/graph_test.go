package graph

import "testing"

func TestGetShortestPathWeight_WithValidInput_ShouldReturnCorrectWeight(t *testing.T) {
	in := [][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
		{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
	}

	g := New(in)
	w := g.GetLowestRiskPath()

	exp := 40

	if exp != w {
		t.Errorf("Incorrect shortest path weight. Expected %d, got %d", exp, w)
	}
}

func TestGetShortestPathWeight_WithTrivialExample_ShouldReturnCorrectWeight(t *testing.T) {
	in := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	g := New(in)
	w := g.GetLowestRiskPath()

	exp := 4

	if exp != w {
		t.Errorf("Incorrect shortest path weight. Expected %d, got %d", exp, w)
	}
}

func TestGetShortestPathWeight_WithNonOneExample_ShouldReturnCorrectWeight(t *testing.T) {
	in := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	g := New(in)
	w := g.GetLowestRiskPath()

	exp := 8

	if exp != w {
		t.Errorf("Incorrect shortest path weight. Expected %d, got %d", exp, w)
	}
}

func TestGetShortestPathWeight_WithBasicExample_ShouldReturnCorrectWeight(t *testing.T) {
	in := [][]int{
		{1, 2, 2},
		{1, 1, 1},
		{2, 2, 1},
	}

	g := New(in)
	w := g.GetLowestRiskPath()

	exp := 4

	if exp != w {
		t.Errorf("Incorrect shortest path weight. Expected %d, got %d", exp, w)
	}
}

func TestGetShortestPathWeight_WithShorterPathUpwards_ShouldUseUpwardsPath(t *testing.T) {
	in := [][]int{
		{1, 100, 1, 1, 1},
		{1, 1, 1, 100, 1},
		{100, 100, 100, 1, 1},
	}

	g := New(in)
	w := g.GetLowestRiskPath()

	exp := 8

	if exp != w {
		t.Errorf("Incorrect shortest path weight. Expected %d, got %d", exp, w)
	}
}
