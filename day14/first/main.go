package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	temp, m := parseInput(in)

	for i := 0; i < 10; i++ {
		temp = processInsertionStep(temp, m)
	}

	score := getScore(temp)

	fmt.Printf("The subtracted quantity is %d\n", score)
}

func getScore(temp string) int {
	leastCom := math.MaxInt
	mostCom := math.MinInt

	count := map[string]int{}

	for i := 0; i < len(temp); i++ {
		r := string(temp[i])

		count[r]++
	}

	for _, ent := range count {
		if ent > mostCom {
			mostCom = ent
		}

		if ent < leastCom {
			leastCom = ent
		}
	}

	return mostCom - leastCom
}

func processInsertionStep(temp string, m map[string]string) string {
	newTemp := ""

	// create the new template while looping over it
	for i := 0; i < len(temp)-1; i++ {
		p := temp[i : i+2]

		newTemp += string(p[0]) + m[p]
	}

	newTemp += string(temp[len(temp)-1])

	return newTemp
}

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
