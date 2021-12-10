package main

import (
	"errors"
	"fmt"
	"sort"

	"github.com/advent-of-code-2021/input"
)

var (
	charMap = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}
)

var (
	ErrCorruptedLine = errors.New("Line contains corrupted syntax")
)

func main() {
	in := input.Read("input.txt")

	score := []int{}

	for _, l := range in {
		lineScore, err := getIncompleteLineScore(l)

		if err != nil {
			continue
		}

		score = append(score, lineScore)
	}

	sort.Ints(score)

	// get the center score
	cen := len(score) / 2

	fmt.Printf("Total autocomplete score is %d\n", score[cen])
}

func getIncompleteLineScore(l string) (int, error) {
	stack := []string{}

	for _, r := range l {
		strr := string(r)
		switch strr {
		case "(", "[", "{", "<":
			// opening char, push to stack
			stack = append(stack, strr)
		case ")", "]", "}", ">":
			if isMatch(stack[len(stack)-1], strr) {
				// we have a matching close tag, pop it
				stack = stack[:len(stack)-1]
			} else {
				// this is a mismatch, meaning a corrupted line. Do nothing with it
				return -1, ErrCorruptedLine
			}
		}
	}

	score := 0
	// reaching this part means the line is not corrupted, check if it's complete
	if len(stack) > 0 {
		// we have a stack left, so let's finish the line and calculate the score
		// run from back to front, otherwise the score will not add up
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			switch stack[i] {
			case "(":
				score += 1
			case "[":
				score += 2
			case "{":
				score += 3
			case "<":
				score += 4
			}
		}
	}

	return score, nil
}

func isMatch(char1, char2 string) bool {
	return charMap[char1] == char2
}
