package main

import "testing"

func TestModelGrowthRate_With18Days_ShouldHaveCorrectFishNum(t *testing.T) {
	in := []string{"3,4,3,1,2"}

	numProcd := ModelGrowthRate(18, in)

	var exp int64 = 26

	if numProcd != exp {
		t.Errorf("Unexpected number of fish processed. Expected %d, got %d", exp, numProcd)
	}
}

func TestModelGrowthRate_With80Days_ShouldHaveCorrectFishNum(t *testing.T) {
	in := []string{"3,4,3,1,2"}

	numProcd := ModelGrowthRate(80, in)

	var exp int64 = 5934

	if numProcd != exp {
		t.Errorf("Unexpected number of fish processed. Expected %d, got %d", exp, numProcd)
	}
}
