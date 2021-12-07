package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/input"
	"github.com/advent-of-code-2021/mathutil"
)

func main() {
	in := input.Read("input.txt")

	spIn := strings.Split(in[0], ",")

	numIn := []int{}
	for _, el := range spIn {
		num, _ := strconv.Atoi(el)

		numIn = append(numIn, num)
	}

	fuel := CalculateFuel(numIn)

	fmt.Printf("Fuel needed to align %d\n", fuel)
}

func CalculateFuel(in []int) int {
	medFloat, _ := mathutil.Median(in)
	med := int(medFloat)

	fuel := 0

	for _, el := range in {
		if el > med {
			fuel += el - med
		} else {
			fuel += med - el
		}
	}

	return fuel
}
