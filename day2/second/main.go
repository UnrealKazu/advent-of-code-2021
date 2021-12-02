package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/input"
)

type Location struct {
	Horizontal int
	Depth      int
	Aim        int
}

func main() {
	in := input.Read("input.txt")

	loc := Location{
		Horizontal: 0,
		Depth:      0,
		Aim:        0,
	}

	for _, i := range in {
		split := strings.Split(i, " ")
		dir := split[0]
		val, _ := strconv.Atoi(split[1])

		loc.Move(dir, val)
	}

	fmt.Printf("Final multiplication value: %d\n", loc.Horizontal*loc.Depth)
}

func (loc *Location) Move(dir string, val int) {
	switch dir {
	case "forward":
		loc.Horizontal += val
		loc.Depth += loc.Aim * val
	case "up":
		loc.Aim -= val
	case "down":
		loc.Aim += val
	}
}
