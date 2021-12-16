package main

import (
	"fmt"
	"testing"
)

func TestHexToBitString_WithSmallString_ShouldReturnCorrectString(t *testing.T) {
	s := "D2FE28"

	bits := HexToBitString(s)

	exp := "110100101111111000101000"

	if bits != exp {
		t.Errorf("Incorrect bitstring. Expected %s, got %s", exp, bits)
	}
}

func TestStupidHexToBitString_WithSmallString_ShouldReturnCorrectString(t *testing.T) {
	s := "D2FE28"

	bits := StupidHexToBitString(s)

	exp := "110100101111111000101000"

	if bits != exp {
		t.Errorf("Incorrect bitstring. Expected %s, got %s", exp, bits)
	}
}

func TestHexToBitString_WithMediumString_ShouldReturnCorrectString(t *testing.T) {
	s := "38006F45291200"

	bits := HexToBitString(s)

	exp := "00111000000000000110111101000101001010010001001000000000"

	if bits != exp {
		t.Errorf("Incorrect bitstring. Expected %s, got %s", exp, bits)
	}
}

func TestGetNextStartIndex_WithLiteralValueType_ShouldCorrectlySkipAhead(t *testing.T) {
	s := "110100101111111000101000"

	i := GetNextStartIndex(s, 4, 6)

	exp := 21

	if i != exp {
		t.Errorf("Unexpected skip index. Expected %d, got %d", exp, i)
	}
}

func TestGetNextStartIndex_WithOperatorTypeZero_ShouldCorrectlySkipAhead(t *testing.T) {
	s := "00111000000000000110111101000101001010010001001000000000"

	i := GetNextStartIndex(s, 5, 6)

	exp := 49

	if i != exp {
		t.Errorf("Unexpected skip index. Expected %d, got %d", exp, i)
	}
}

func TestGetNextStartIndex_WithOperatorTypeOne_ShouldCorrectlySkipAhead(t *testing.T) {
	s := "11101110000000001101010000001100100000100011000001100000"

	i := GetNextStartIndex(s, 5, 6)

	exp := 51

	if i != exp {
		t.Errorf("Unexpected skip index. Expected %d, got %d", exp, i)
	}
}

func Test(t *testing.T) {
	s := "8A004A801A8002F478"

	sum := processTransmission(s)

	fmt.Println(sum)
}

func TestProcessTransmission_WithTransmissions_ShouldReturnCorrectSum(t *testing.T) {
	var tests = []struct {
		a   string
		exp int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, tt := range tests {
		t.Run(tt.a, func(t *testing.T) {
			got := processTransmission(tt.a)
			if got != tt.exp {
				t.Errorf("Incorrect sum value. Expected %d, got %d", tt.exp, got)
			}
		})
	}
}
