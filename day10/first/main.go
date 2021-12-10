package main

import (
	"fmt"

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

func main() {
	in := input.Read("input.txt")

	score := 0

	for _, l := range in {
		ill := getIllegalCharacter(l)

		switch ill {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}

	fmt.Printf("Total syntax error score is %d\n", score)
}

func getIllegalCharacter(l string) string {
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
				// this is a mismatch, return it
				return strr
			}
		}
	}

	return ""
}

func isMatch(char1, char2 string) bool {
	return charMap[char1] == char2
}
