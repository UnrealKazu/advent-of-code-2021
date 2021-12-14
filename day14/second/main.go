package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/advent-of-code-2021/input"
)

const (
	MaxStep = 40
)

func main() {
	in := input.Read("input.txt")

	score := startPolymerization(in)

	fmt.Printf("The subtracted quantity is %v\n", score)
}

func startPolymerization(in []string) int64 {
	temp, m := parseInput(in)

	count := map[string]int64{}

	// initialize the count with all possible pairs
	for ent := range m {
		count[ent] = 0
	}

	// increase the pair count for the initial pairs
	for i := 0; i < len(temp)-1; i++ {
		count[temp[i:i+2]]++
	}

	for i := 0; i < MaxStep; i++ {
		shad := map[string]int64{}

		// because we are modifying the count of pairs while we are
		// looping over it, make a shadow copy first, so that we
		// can modify the count independently
		for ent, num := range count {
			shad[ent] = num
		}

		// for each tracked pair, calculate the number of new pairs
		for ent, num := range shad {
			if num > 0 {
				// only process pairs that we actually track
				processPair(ent, num, m, count)
			}
		}
	}

	return getScore(temp, count)
}

// getScore determines the score based on the most and least common character (not pair)
// in the final polymer template
func getScore(temp string, count map[string]int64) int64 {
	charCount := map[string]int64{}

	// get the individual letter counts
	for ent, num := range count {
		if num > 0 {
			// again, only count tracked pairs
			// but we are only interested in the first character
			// (because the second is duplicated in the next pair)
			charCount[string(ent[0])] += num
		}
	}

	// and add one for the final character of the template (because it's left out above)
	charCount[string(temp[len(temp)-1])]++

	// determine the most common and least common character
	var leastCom int64 = math.MaxInt64
	var mostCom int64 = math.MinInt64

	for _, ent := range charCount {
		if ent > mostCom {
			mostCom = ent
		}

		if ent < leastCom {
			leastCom = ent
		}
	}

	return mostCom - leastCom
}

// processPair splits the given pair and tracks the count of both the old and new pairs
func processPair(p string, num int64, m map[string]string, count map[string]int64) {
	ins := m[p]

	// decrease the count of the incoming pair, because we're splitting it up, so losing it
	// times the number of pairs
	count[p] -= num

	np1 := string(p[0]) + ins
	np2 := ins + string(p[1])

	// count the new pairs
	count[np1] += num
	count[np2] += num
}

// parseInput converts the input slice into a tuple of the initial template and a map
// of pair -> inserted character combinations
func parseInput(in []string) (string, map[string]string) {
	temp := in[0]

	m := make(map[string]string, len(in)-2)
	for i := 2; i < len(in); i++ {
		spl := strings.Split(in[i], " -> ")

		// add pair to map, with the insertion as a result
		m[spl[0]] = spl[1]
	}

	return temp, m
}
