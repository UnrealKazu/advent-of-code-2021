package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code-2021/input"
)

const (
	MaxDays    int = 256
	TimerReset int = 7
)

type Fish struct {
	Timer      int
	ObserveDay int
}

func main() {
	in := input.Read("input.txt")

	start := time.Now()

	numProcd := ModelGrowthRate(MaxDays, in)

	fmt.Printf("Calculation done in %s\n", time.Since(start))
	fmt.Printf("Number of fishes after %d days: %d\n", MaxDays, numProcd)
}

func ModelGrowthRate(maxDays int, in []string) int64 {
	initF := strings.Split(in[0], ",")

	// put all fish in the channel
	var num int64 = 0

	for _, f := range initF {
		timer, _ := strconv.Atoi(f)
		ff := Fish{
			Timer:      timer,
			ObserveDay: 0,
		}

		num += Process(maxDays, ff)
	}

	return num
}
