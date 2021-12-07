package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/input"
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
	med := GetMedian(in)

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

func GetMedian(in []int) int {
	sort.Ints(in)

	cen := len(in) / 2

	var med int
	if len(in)%2 == 0 {
		// even num of ints, so take average of the center
		med = (in[cen-1] + in[cen]) / 2
	} else {
		med = in[cen]
	}

	return med
}
