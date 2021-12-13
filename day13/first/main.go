package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/day13/first/paper"
	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	r := regexp.MustCompile("fold along (x|y)=([0-9]+)")

	p := paper.New()

	i := 0
	cur := in[i]

	for cur != "" {
		spl := strings.Split(cur, ",")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])

		p.AddDot(x, y)
		i++
		cur = in[i]
	}

	// get the first fold operation
	i++
	foldOp := in[i]
	op := r.FindStringSubmatch(foldOp)

	dir := op[1]
	fold, _ := strconv.Atoi(op[2])

	if dir == "x" {
		p.FoldX(fold)
	} else {
		p.FoldY(fold)
	}

	nrof := p.GetNrofDots()

	fmt.Printf("The number of dots after one fold is %d\n", nrof)
}
