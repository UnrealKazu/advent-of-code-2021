package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	sum := SumOutputValues(in)

	fmt.Printf("Sum of all output values is %d\n", sum)
}

func SumOutputValues(in []string) int {
	sum := 0

	for _, l := range in {

		spl := strings.Split(l, " | ")

		lang := strings.Split(spl[0], " ")
		nums := strings.Split(spl[1], " ")

		numStr := DetermineNumber(nums, lang)

		num, _ := strconv.Atoi(numStr)

		sum += num
	}

	return sum
}

func DetermineNumber(nums []string, lang []string) string {
	mapping := [10]string{}

	fives := []string{}
	sixes := []string{}

	// first determine 1,4,7,8
	for _, ent := range lang {
		sortNum := SortString(ent)
		switch len(ent) {
		case 2:
			mapping[1] = SortString(sortNum)
		case 4:
			mapping[4] = SortString(sortNum)
		case 3:
			mapping[7] = SortString(sortNum)
		case 7:
			mapping[8] = SortString(sortNum)
		case 5:
			fives = append(fives, SortString(sortNum))
		case 6:
			sixes = append(sixes, SortString(sortNum))
		}
	}

	i := 0
	for len(sixes) > 0 {
		cur := sixes[i]

		// we have no 9, so check for that first
		if mapping[9] == "" {

			// check if it contains a 4
			if ContainsSegments(cur, mapping[4]) {
				// this is a 9
				mapping[9] = cur

				// remove this entry
				sixes = append(sixes[:i], sixes[i+1:]...)
				if len(sixes) > 0 {
					i = (i + 1) % len(sixes)
				}
			} else {
				if len(sixes) > 0 {
					i = (i + 1) % len(sixes)
				}
			}
		} else {
			// check if it's a 0 or a 6
			if ContainsSegments(cur, mapping[1]) {
				// only a 0 can contain a 1
				mapping[0] = cur

				// remove this entry
				sixes = append(sixes[:i], sixes[i+1:]...)
				if len(sixes) > 0 {
					i = (i + 1) % len(sixes)
				}
			} else {
				// this must be a 6
				mapping[6] = cur

				// remove this entry
				sixes = append(sixes[:i], sixes[i+1:]...)
				if len(sixes) > 0 {
					i = (i + 1) % len(sixes)
				}
			}
		}
	}

	for len(fives) > 0 {
		cur := fives[0]

		if ContainsSegments(mapping[9], cur) {
			// only a 3 and a 5 are contained in a 9
			if ContainsSegments(cur, mapping[1]) {
				// only a 3 contains a 1
				mapping[3] = cur

				// remove this entry
				fives = append(fives[:i], fives[i+1:]...)
			} else {
				// the only remaining case, this must be a 5
				mapping[5] = cur

				// remove this entry
				fives = append(fives[:i], fives[i+1:]...)
			}
		} else {
			mapping[2] = cur

			// remove this entry
			fives = append(fives[:i], fives[i+1:]...)
		}
	}

	// invert the mapping
	invMap := make(map[string]int, 10)

	for k, m := range mapping {
		invMap[m] = k
	}

	// now, finally loop over the numbers
	numStr := ""

	for _, num := range nums {
		sortNum := SortString(num)
		numStr += strconv.Itoa(invMap[sortNum])
	}

	return numStr
}

func ContainsSegments(base string, sub string) bool {
	for _, r := range sub {
		if !strings.Contains(base, string(r)) {
			return false
		}
	}

	return true
}

func ReplaceRunes(base string, sub string) string {
	for _, s := range sub {
		base = strings.ReplaceAll(base, string(s), "")
	}

	return base
}
