package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code-2021/input"
)

func main() {
	start := time.Now()
	in := input.Read("input.txt")

	spIn := strings.Split(in[0], ",")

	numIn := []int{}
	for _, el := range spIn {
		num, _ := strconv.Atoi(el)

		numIn = append(numIn, num)
	}

	fuel := DetermineLeastFuel(numIn)

	fmt.Printf("Fuel needed to align %d\n", fuel)
	fmt.Printf("Calculation time %s\n", time.Since(start))
}

func DetermineLeastFuel(in []int) int {
	mean := GetMean(in)
	flMean := int(math.Floor(float64(mean)))
	ceilMean := int(math.Ceil(float64(mean)))

	flFuel := CalculateFuel(flMean, in)
	ceilFuel := CalculateFuel(ceilMean, in)

	if flFuel > ceilFuel {
		return ceilFuel
	} else {
		return flFuel
	}
}

func CalculateFuel(mean int, in []int) int {
	fuel := 0

	for _, el := range in {
		var diff int
		if el > mean {
			diff += el - mean
		} else {
			diff += mean - el
		}

		fuel += (diff * (diff + 1) / 2)
	}

	return fuel
}

func GetMean(in []int) float64 {
	sum := 0

	for _, el := range in {
		sum += el
	}

	return float64(sum) / float64(len(in))
}
