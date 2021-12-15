package main

import (
	"reflect"
	"testing"
)

func TestMultiplyMap_WithHorizontalMultiply_ShouldProduceCorrectMap(t *testing.T) {
	m := [][]int{
		{8},
	}

	mm := multiplyMap(m, 5, 1)

	exp := [][]int{
		{8, 9, 1, 2, 3},
	}

	if !reflect.DeepEqual(exp, mm) {
		t.Errorf("Incorrect multiplied map. Expected %v, got %v", exp, mm)
	}
}

func TestMultiplyMap_WithHorizontalMultiplyAndTwoRows_ShouldProduceCorrectMap(t *testing.T) {
	m := [][]int{
		{8},
		{9},
	}

	mm := multiplyMap(m, 5, 1)

	exp := [][]int{
		{8, 9, 1, 2, 3},
		{9, 1, 2, 3, 4},
	}

	if !reflect.DeepEqual(exp, mm) {
		t.Errorf("Incorrect multiplied map. Expected %v, got %v", exp, mm)
	}
}

func TestMultiplyMap_WithVerticalMultiply_ShouldProduceCorrectMap(t *testing.T) {
	m := [][]int{
		{8},
	}

	mm := multiplyMap(m, 1, 5)

	exp := [][]int{
		{8},
		{9},
		{1},
		{2},
		{3},
	}

	if !reflect.DeepEqual(exp, mm) {
		t.Errorf("Incorrect multiplied map. Expected %v, got %v", exp, mm)
	}
}

func TestMultiplyMap_WithHorizontalAndVertical_ShouldProduceCorrectMap(t *testing.T) {
	m := [][]int{
		{8},
	}

	mm := multiplyMap(m, 5, 5)

	exp := [][]int{
		{8, 9, 1, 2, 3},
		{9, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
	}

	if !reflect.DeepEqual(exp, mm) {
		t.Errorf("Incorrect multiplied map. Expected %v, got %v", exp, mm)
	}
}
