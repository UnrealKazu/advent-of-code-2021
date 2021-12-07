package mathutil

import "testing"

func TestMean_WithPositiveInput_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{7, 2, 3}

	mean, err := Mean(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := 4.0

	if mean != exp {
		t.Errorf("Unexpected mean value. Expected %f, got %f", exp, mean)
	}
}

func TestMean_WithNegativeInput_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-7, -2, -3}

	mean, err := Mean(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := -4.0

	if mean != exp {
		t.Errorf("Unexpected mean value. Expected %f, got %f", exp, mean)
	}
}

func TestMean_WithPositiveAndNegativeInput_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-8, 2, 3}

	mean, err := Mean(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := -1.0

	if mean != exp {
		t.Errorf("Unexpected mean value. Expected %f, got %f", exp, mean)
	}
}

func TestMean_WithEmptyInput_ShouldReturnError(t *testing.T) {
	in := []int{}

	mean, err := Mean(in)

	exp := float64(-1)

	if mean != -1 {
		t.Errorf("Unexpected mean return. Expected %f, got %f", exp, mean)
	}

	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestMedian_WithPositiveInputOdd_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{7, 2, 3}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := 3.0

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithNegativeInputOdd_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-7, -2, -3}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := -3.0

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithPositiveAndNegativeInputOdd_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-8, 2, 3}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := 2.0

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithPositiveInputEven_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{7, 2, 3, 1}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := 2.5

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithNegativeInputEven_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-7, -2, -3, -1}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := -2.5

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithPositiveAndNegativeInputEven_ShouldReturnCorrectValue(t *testing.T) {
	in := []int{-8, 2, 3, 1}

	median, err := Median(in)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	exp := 1.5

	if median != exp {
		t.Errorf("Unexpected median value. Expected %f, got %f", exp, median)
	}
}

func TestMedian_WithEmptyInput_ShouldReturnError(t *testing.T) {
	in := []int{}

	median, err := Median(in)

	exp := float64(-1)

	if median != -1 {
		t.Errorf("Unexpected median return. Expected %f, got %f", exp, median)
	}

	if err == nil {
		t.Error("Expected error, got none")
	}
}
