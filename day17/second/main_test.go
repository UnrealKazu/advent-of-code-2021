package main

import "testing"

func TestFindHitCount_WithExampe_ShouldReturnCorrectValue(t *testing.T) {
	tt := NewTarget(20, 30, -5, -10)

	y := findHitCount(tt)

	exp := 112

	if y != exp {
		t.Errorf("Incorrect y-value. Expected %d, got %d", exp, y)
	}
}

func TestHitsTarget_WithValidVector_ShouldHitTarget(t *testing.T) {
	tt := NewTarget(20, 30, -5, -10)

	hit := tt.hitsTarget(6, 3)

	exp := true

	if hit != exp {
		t.Errorf("Incorrect hit. Expected %v, got %v", exp, hit)
	}
}
