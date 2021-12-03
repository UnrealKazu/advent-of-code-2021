package main

import "testing"

func TestCalculateRates_WithFullRecord_ShouldReturnCorrectRates(t *testing.T) {
	rec := make([]int, 4)
	rec[0] = 1
	rec[1] = 1
	rec[2] = 1
	rec[3] = 1

	gamma, epsilon := CalculateRates(rec)

	if gamma != 15 {
		t.Errorf("Unexpected gamma rate. Expected %d, got %d", 15, gamma)
	}

	if epsilon != 0 {
		t.Errorf("Unexpected epsilon rate. Expected %d, got %d", 0, epsilon)
	}
}

func TestCalculateRates_WithEmptyRecord_ShouldReturnCorrectRates(t *testing.T) {
	rec := make([]int, 4)
	rec[0] = -1
	rec[1] = -1
	rec[2] = -1
	rec[3] = -1

	gamma, epsilon := CalculateRates(rec)

	if gamma != 0 {
		t.Errorf("Unexpected gamma rate. Expected %d, got %d", 0, gamma)
	}

	if epsilon != 15 {
		t.Errorf("Unexpected epsilon rate. Expected %d, got %d", 15, epsilon)
	}
}

func TestCalculateRates_WithMixedRecord_ShouldReturnCorrectRates(t *testing.T) {
	rec := make([]int, 4)
	rec[0] = 1
	rec[1] = 1
	rec[2] = -1
	rec[3] = -1

	gamma, epsilon := CalculateRates(rec)

	if gamma != 12 {
		t.Errorf("Unexpected gamma rate. Expected %d, got %d", 12, gamma)
	}

	if epsilon != 3 {
		t.Errorf("Unexpected epsilon rate. Expected %d, got %d", 3, epsilon)
	}
}
