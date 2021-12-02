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
}

func main() {
	in := input.Read("input.txt")

	loc := Location{
		Horizontal: 0,
		Depth:      0,
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
	case "up":
		loc.Depth -= val
	case "down":
		loc.Depth += val
	}
}
