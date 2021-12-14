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

func TestProcessInsertionStep_WithValidTemplateAndMap_ShouldProduceCorrectTemplate(t *testing.T) {
	temp := "NNCB"

	m := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
		"HB": "C",
		"HC": "B",
		"HN": "C",
		"NN": "C",
		"BH": "H",
		"NC": "B",
		"NB": "B",
		"BN": "B",
		"BB": "N",
		"BC": "B",
		"CC": "N",
		"CN": "C",
	}

	ntemp := processInsertionStep(temp, m)

	exp := "NCNBCHB"

	if ntemp != exp {
		t.Errorf("Incorrect template after step. Expected %s, got %s", exp, ntemp)
	}
}
