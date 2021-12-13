package main

import "testing"

func TestGetNrofSmallCavePaths_WithSmallInputSet_ShouldReturnCorrectValue(t *testing.T) {
	in := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}

	num := GetNrofSmallCavePaths(in)

	exp := 10

	if num != exp {
		t.Errorf("Number of paths is incorrect. Expected %d, got %d\n", exp, num)
	}
}

func TestGetNrofSmallCavePaths_WithMediumInputSet_ShouldReturnCorrectValue(t *testing.T) {
	in := []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}

	num := GetNrofSmallCavePaths(in)

	exp := 19

	if num != exp {
		t.Errorf("Number of paths is incorrect. Expected %d, got %d\n", exp, num)
	}
}

func TestGetNrofSmallCavePaths_WithLargeInputSet_ShouldReturnCorrectValue(t *testing.T) {
	in := []string{
		"fs-end",
		"he-DX",
		"fs-he",
		"start-DX",
		"pj-DX",
		"end-zg",
		"zg-sl",
		"zg-pj",
		"pj-he",
		"RW-he",
		"fs-DX",
		"pj-RW",
		"zg-RW",
		"start-pj",
		"he-WI",
		"zg-he",
		"pj-fs",
		"start-RW",
	}

	num := GetNrofSmallCavePaths(in)

	exp := 226

	if num != exp {
		t.Errorf("Number of paths is incorrect. Expected %d, got %d\n", exp, num)
	}
}
