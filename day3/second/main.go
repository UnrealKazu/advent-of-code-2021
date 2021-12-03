package main

import (
	"fmt"

	"github.com/advent-of-code-2021/input"
)

type RatingType int

const (
	Oxygen RatingType = iota
	CO2
)

func main() {
	in1 := input.Read("input.txt")
	in2 := make([]string, len(in1))
	copy(in2, in1)

	oxy := GetRating(Oxygen, in1)
	co2 := GetRating(CO2, in2)

	fmt.Printf("Submarine life support rating is %d\n", oxy*co2)
}

func GetRating(rateType RatingType, report []string) uint {
	byteString := ""
	for i := 0; i < len(report[0]); i++ {
		rec := 0
		for _, bit := range report {
			// check if bit is 0 (increase record) or 1 (decrease record)
			if bit[i] == '0' {
				rec--
			} else {
				// bit is 1
				rec++
			}
		}

		if rec >= 0 {
			if rateType == Oxygen {
				byteString += "1"
			} else if rateType == CO2 {
				byteString += "0"
			}
		} else {
			if rateType == Oxygen {
				byteString += "0"
			} else if rateType == CO2 {
				byteString += "1"
			}
		}

		// trim the report. Remove all entries that do (partially)
		// start with our current byteString
		filteredRep := make([]string, 0)
		for _, entry := range report {
			if entry[:len(byteString)] == byteString {
				filteredRep = append(filteredRep, entry)
			}
		}
		report = filteredRep

		if len(report) == 1 {
			// found the final rating, break
			break
		}
	}

	return CalculateRateValue(report[0])
}

func CalculateRateValue(record string) uint {
	var rate uint = 0
	for i := 0; i < len(record); i++ {
		// perform bitshifts based on the value
		if record[i] == '1' {
			rate = rate | (1 << (len(record) - 1 - i))
		}
	}

	return rate
}
