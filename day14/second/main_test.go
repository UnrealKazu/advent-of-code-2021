package main

import (
	"reflect"
	"testing"
)

func TestParseInput_WithValidInput_ShouldReturnCorrectPair(t *testing.T) {
	in := []string{
		"NNCB",
		"",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
	}

	temp, m := parseInput(in)

	tempExp := "NNCB"
	mExp := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
	}

	if temp != tempExp {
		t.Errorf("Incorrect template. Expected %s, got %s", tempExp, temp)
	}

	if !reflect.DeepEqual(m, mExp) {
		t.Errorf("Incorrect pair map. Expected %v, got %v", mExp, m)
	}
}
