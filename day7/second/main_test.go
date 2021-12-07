package main

import "testing"

func TestCalculateFuel_WithValidInput_ShouldGetCorrectValue(t *testing.T) {
	in := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := DetermineLeastFuel(in)

	exp := 168

	if fuel != exp {
		t.Errorf("Unexpected fuel number. Expected %d, got %d", exp, fuel)
	}
}
