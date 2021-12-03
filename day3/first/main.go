package main

import (
	"fmt"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	// make a slice based on the length of the first row (since all rows are equal in size)
	rec := make([]int, len(in[0]))

	for _, bit := range in {
		for i := 0; i < len(bit); i++ {
			// check if bit is 0 (increase record) or 1 (decrease record)
			if bit[i] == '0' {
				rec[i]--
			} else {
				// bit is 1
				rec[i]++
			}
		}
	}

	gamma, epsilon := CalculateRates(rec)

	fmt.Printf("Submarine power consumption is %d\n", gamma*epsilon)
}

func CalculateRates(record []int) (uint, uint) {
	var gamma uint = 0
	var epsilon uint = 0
	for i := 0; i < len(record); i++ {
		// perform bitshifts based on the most common or least common bit
		if record[i] > 0 {
			gamma = gamma | (1 << (len(record) - 1 - i))
		} else {
			epsilon = epsilon | (1 << (len(record) - 1 - i))
		}
	}

	return gamma, epsilon
}
